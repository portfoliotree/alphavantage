// Package alphavantage implements clients and parsers for the
// https://www.alphavantage.co API.

package alphavantage

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

const (
	StandardTokenEnvironmentVariableName = "ALPHA_VANTAGE_TOKEN"
)

const DefaultDateFormat = "2006-01-02"

type Client struct {
	Limiter interface {
		Wait(ctx context.Context) error
	}
	Client interface {
		Do(*http.Request) (*http.Response, error)
	}
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		Client:  http.DefaultClient,
		Limiter: rate.NewLimiter(rate.Every(time.Minute/5), 5),
		APIKey:  apiKey,
	}
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	if client.Limiter != nil {
		err := client.Limiter.Wait(req.Context())
		if err != nil {
			return &http.Response{}, err
		}
	}

	q := req.URL.Query()
	q.Set("apikey", client.APIKey)
	req.URL.RawQuery = q.Encode()

	res, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		buf, err := io.ReadAll(io.LimitReader(res.Body, 1<<10))
		if err != nil {
			buf = []byte(err.Error())
		}
		return res, fmt.Errorf("request failed with status %d %s: %s",
			res.StatusCode, http.StatusText(res.StatusCode), string(buf))
	}

	return res, nil
}

func checkError(rc io.ReadCloser) (io.ReadCloser, error) {
	var buf [1]byte
	n, err := rc.Read(buf[:])
	if err != nil {
		return nil, fmt.Errorf("could not read request response: %w", err)
	}

	mr := io.MultiReader(bytes.NewReader(buf[:]), rc)
	if n > 0 && buf[0] == '{' {
		var message struct {
			Note         string `json:"Note,omitempty"`
			Information  string `json:"Information,omitempty"`
			ErrorMessage string `json:"Error Message,omitempty"`
			Detail       string `json:"detail,omitempty"`
		}
		err = json.NewDecoder(mr).Decode(&message)
		if err != nil {
			return nil, fmt.Errorf("could not read response for: %w", err)
		}
		if strings.Contains(message.Note, " higher API call frequency") {
			return nil, fmt.Errorf("reached alphavantage rate limit")
		}

		if message.ErrorMessage != "" {
			return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(message.ErrorMessage))
		}
		if message.Detail != "" {
			return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(message.Detail))
		}
		if message.Note != "" || message.Information != "" {
			return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(strings.Join([]string{message.Note, message.Information}, " ")))
		}

		return nil, fmt.Errorf("alphavantage request did not return csv")
	}

	return multiReadCloser{
		Reader: mr,
		close:  rc.Close,
	}, nil
}

var typeType = reflect.TypeOf(time.Time{})

// ParseCSV parses rows of data into a slice of structs it only supports decoding into
// fields of type string, int, float64, and time.Time
//
// Struct fields must be tagged with their expected column header with "column-name".
// If there is no mapping for a column to field, that data column is ignored and the struct
// field will remain unset (it will have its zero value).
//
// Fields with type time.Time may use an additional "time-layout" field
// to specify the layout to use with time.ParseInLocation.
// If location is not specified, eastern time is used.
// "null" values in CSV for time are ignored; time keeps its zero value.
func ParseCSV[T any](r io.Reader, data *[]T, location *time.Location) error {
	if data == nil {
		panic(fmt.Errorf("data must not be nil"))
	}
	var err error
	for row := range ParseCSVRows[T](r, location, func(e error) bool {
		err = e
		return false
	}) {
		*data = append(*data, row)
	}
	return err
}

func ParseCSVRows[T any](r io.Reader, location *time.Location, handleErr func(error) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		if location == nil {
			location = time.UTC
		}

		rowType := reflect.TypeFor[T]()

		reader := csv.NewReader(bufio.NewReader(r))
		reader.TrimLeadingSpace = true
		header, err := reader.Read()
		if err != nil {
			handleErr(err)
			return
		}
		reader.FieldsPerRecord = len(header)

		if rowType.Kind() != reflect.Struct {
			panic(fmt.Errorf("expected a struct kind: got %s", rowType.Kind()))
			return
		}

		structType := rowType

		columnToField := make(map[int]int, len(header))
		for columnHeaderIndex, columnHeaderName := range header {
			for fieldIndex := 0; fieldIndex < structType.NumField(); fieldIndex++ {
				fieldType := structType.Field(fieldIndex)

				csvTag := fieldType.Tag.Get("column-name")
				if csvTag != columnHeaderName {
					continue
				}

				columnToField[columnHeaderIndex] = fieldIndex
			}
		}

		for rowIndex := 1; ; rowIndex++ {
			row, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					return
				}
				handleErr(err)
				return
			}

			structValue := reflect.New(structType)

			for columnIndex, value := range row {
				fieldIndex, ok := columnToField[columnIndex]
				if !ok {
					continue
				}

				structFieldType := structType.Field(fieldIndex)

				switch structFieldType.Type.Kind() {
				case reflect.String:
					structValue.Elem().Field(fieldIndex).SetString(value)
				case reflect.Float64:
					fl, err := strconv.ParseFloat(value, 64)
					if err != nil {
						handleErr(fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w", value, rowIndex, columnIndex, header[columnIndex], err))
						return
					}
					structValue.Elem().Field(fieldIndex).SetFloat(fl)
				case reflect.Int:
					in, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						handleErr(fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w", value, rowIndex, columnIndex, header[columnIndex], err))
						return
					}
					structValue.Elem().Field(fieldIndex).SetInt(in)
				default:
					if structFieldType.Type != typeType {
						handleErr(fmt.Errorf("unsupported type %T for field %s", structFieldType.Type, structFieldType.Name))
						return
					}

					layout := DefaultDateFormat
					tagLayout := structFieldType.Tag.Get("time-layout")
					if tagLayout != "" {
						layout = tagLayout
					}
					if value == "null" {
						continue
					}
					tm, err := time.ParseInLocation(layout, value, location)
					if err != nil {
						handleErr(fmt.Errorf("failed to parse time value on row %d column %d (%s): %w", rowIndex, columnIndex, header[columnIndex], err))
						return
					}
					structValue.Elem().Field(fieldIndex).Set(reflect.ValueOf(tm))
				}
			}

			if !yield(structValue.Elem().Interface().(T)) {
				return
			}
		}
	}
}
