package specification

import (
	"encoding/json"
	"fmt"
	"slices"
)

const JSONIndent = "\t"

type QueryParameter struct {
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	Values    []EnumValue               `json:"values,omitempty"`
	ValuesURL string                    `json:"values_url,omitempty"`
	Validate  []QueryParameterValidator `json:"validate,omitempty"`
	Format    string                    `json:"format,omitempty"`
}

type EnumValue struct {
	json.RawMessage
}

func (v EnumValue) Name() (string, error) {
	switch v.RawMessage[0] {
	case '{':
		var elem struct {
			Value string `json:"value"`
		}
		if err := json.Unmarshal([]byte(v.RawMessage), &elem); err != nil {
			return "", err
		}
		return elem.Value, nil
	case '"':
		var str string
		if err := json.Unmarshal([]byte(v.RawMessage), &str); err != nil {
			return "", err
		}
		return str, nil
	default:
		return "", fmt.Errorf("unexpected query parameter value shape `%s`", string(v.RawMessage))
	}
}

type QueryParameterValidator struct {
	Message    string `json:"message"`
	Identifier string `json:"identifier"`
	Expression string `json:"expression"`
}

type Generate struct {
	Files []File `json:"files"`
}

type File struct {
	File      string     `json:"file"`
	Functions []Function `json:"functions"`
}

type CSVColumn struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Format string `json:"format,omitempty"`
}

type Function struct {
	Name       string              `json:"name"`
	Required   []string            `json:"required"`
	Optional   []string            `json:"optional"`
	EnumSubset map[string][]string `json:"enum_subset,omitempty"`
	Examples   []string            `json:"examples"`
	CSVColumns []CSVColumn         `json:"csv_columns,omitempty"`
}

func (fn Function) HasDatatypeParameter() bool {
	return slices.Contains(fn.Optional, QueryKeyDataType) || slices.Contains(fn.Required, QueryKeyDataType)
}
