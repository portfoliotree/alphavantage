package alphavantage

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

var easternTimezone *time.Location

func init() {
	var err error
	easternTimezone, err = time.LoadLocation("US/Eastern")
	if err != nil {
		panic(err)
	}
}

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
		Limiter: rate.NewLimiter(rate.Every(time.Minute/5), 2),
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
		buf, err := ioutil.ReadAll(io.LimitReader(res.Body, 1<<10))
		if err != nil {
			buf = []byte(err.Error())
		}
		return res, fmt.Errorf("request failed with status %d %s: %s",
			res.StatusCode, http.StatusText(res.StatusCode), string(buf))
	}

	return res, nil
}

func checkError(r io.Reader) (io.Reader, error) {
	var buf [1]byte
	n, err := r.Read(buf[:])
	if err != nil {
		return nil, fmt.Errorf("could not read request response: %w", err)
	}

	mr := io.MultiReader(bytes.NewReader(buf[:]), r)
	if n > 0 && buf[0] == '{' {
		var message struct {
			Note        string `json:"Note"`
			Information string `json:"Information"`
		}
		err = json.NewDecoder(mr).Decode(&message)
		if err != nil {
			return nil, fmt.Errorf("could not read response for: %w", err)
		}
		if strings.Contains(message.Note, " higher API call frequency") {
			return nil, fmt.Errorf("reached alphavantage rate limit")
		}

		return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(strings.Join([]string{message.Note, message.Information}, " ")))
	}

	return mr, nil
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
func ParseCSV(r io.Reader, data interface{}, location *time.Location) error {
	if location == nil {
		location = easternTimezone
	}

	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr {
		panic("parse must receive pointer to data")
	}
	if rv.IsNil() {
		panic("parse must not receive pointer to nil data")
	}

	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true
	header, err := reader.Read()
	if err != nil {
		return err
	}
	reader.FieldsPerRecord = len(header)

	rvt := rv.Type()
	switch rvt.Elem().Kind() {
	default:
		return fmt.Errorf("expected a pointer to an array or slice: got %T", data)
	case reflect.Slice:
	}
	if rvt.Elem().Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to an array or slice of structs: got %T", data)
	}

	columnToField := make(map[int]int, len(header))
	structType := rvt.Elem().Elem()
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
				return nil
			}
			return err
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
					return fmt.Errorf("failed to parse float64 value %q on row %d column %d: %w", value, rowIndex, columnIndex, err)
				}
				structValue.Elem().Field(fieldIndex).SetFloat(fl)
			case reflect.Int:
				in, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to parse float64 value %q on row %d column %d: %w", value, rowIndex, columnIndex, err)
				}
				structValue.Elem().Field(fieldIndex).SetInt(in)
			default:
				if structFieldType.Type != typeType {
					return fmt.Errorf("unsupported type %T for field %s", structFieldType.Type, structFieldType.Name)
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
					return fmt.Errorf("failed to parse time value %q on row %d column %d: %w", value, rowIndex, columnIndex, err)
				}
				structValue.Elem().Field(fieldIndex).Set(reflect.ValueOf(tm))
			}
		}

		rv.Elem().Set(reflect.Append(rv.Elem(), structValue.Elem()))
	}

	return nil
}
