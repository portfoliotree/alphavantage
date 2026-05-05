package api

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// DefaultDateFormat is the RFC 3339 date format used for parsing dates.
const DefaultDateFormat = "2006-01-02"

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

type multiReadCloser struct {
	io.Reader
	close func() error
}

func (mrc multiReadCloser) Close() error {
	return mrc.close()
}

// ParseCSV parses CSV data into a slice of structs using reflection.
//
// Supported field types:
//   - string: Direct mapping from CSV column value
//   - int: Parsed using strconv.ParseInt with base 10
//   - float64: Parsed using strconv.ParseFloat
//   - time.Time: Parsed using time.ParseInLocation (see time-layout tag)
//
// Struct field tags:
//   - `column-name:"header"`: Maps field to CSV column header (required)
//   - `time-layout:"layout"`: Custom time format for time.Time fields (optional, defaults to "2006-01-02")
//
// Example struct:
//
//	type StockPrice struct {
//	    Date   time.Time `column-name:"timestamp"`
//	    Open   float64   `column-name:"open"`
//	    High   float64   `column-name:"high"`
//	    Volume int       `column-name:"volume"`
//	}
//
// Unmapped columns are ignored. Fields without matching columns keep their zero value.
// Time fields with "null" values remain as zero time.Time.
func ParseCSV[T any](r io.Reader, data *[]T, location *time.Location) error {
	if data == nil {
		panic(fmt.Errorf("data must not be nil"))
	}
	var err error
	for row := range parseCSVRows[T](ensureReadCloser(r), location, func(e error) bool {
		err = e
		return false
	}) {
		*data = append(*data, row)
	}
	return err
}

func ensureReadCloser(r io.Reader) io.ReadCloser {
	if rc, ok := r.(io.ReadCloser); ok {
		return rc
	}
	return io.NopCloser(r)
}

// parseCSVRows returns an iterator that parses CSV data row by row into structs.
// This is memory-efficient for large datasets as it processes one row at a time.
//
// Uses the same struct field tagging system as ParseCSV:
//   - `column-name:"header"`: Maps field to CSV column header (required)
//   - `time-layout:"layout"`: Custom time format for time.Time fields (optional)
//
// The handleErr function is called when parsing errors occur. Return true to continue
// processing, false to stop. Location defaults to UTC if nil.
//
// Example usage:
//
//	for price := range parseCSVRows[StockPrice](reader, time.UTC, func(err error) bool {
//	    log.Printf("Parse error: %v", err)
//	    return true // continue on errors
//	}) {
//	    fmt.Printf("Price: %+v\n", price)
//	}
func parseCSVRows[T any](r io.Reader, location *time.Location, handleErr func(error) bool) iter.Seq[T] {
	rc := ensureReadCloser(r)
	return func(yield func(T) bool) {
		defer func() { _ = rc.Close() }()
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
						if handleErr(fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w", value, rowIndex, columnIndex, header[columnIndex], err)) {
							continue
						}
						return
					}
					structValue.Elem().Field(fieldIndex).SetFloat(fl)
				case reflect.Int:
					in, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						if !handleErr(fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w", value, rowIndex, columnIndex, header[columnIndex], err)) {
							continue
						}
						return
					}
					structValue.Elem().Field(fieldIndex).SetInt(in)
				default:
					if structFieldType.Type != typeType {
						if handleErr(fmt.Errorf("unsupported type %T for field %s", structFieldType.Type, structFieldType.Name)) {
							continue
						}
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
						if handleErr(fmt.Errorf("failed to parse time value on row %d column %d (%s): %w", rowIndex, columnIndex, header[columnIndex], err)) {
							continue
						}
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
