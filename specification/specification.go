package specification

import (
	"encoding/json"
)

type QueryParameter struct {
	Ident     string                    `json:"ident,omitempty"`
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	Values    []json.RawMessage         `json:"values,omitempty"`
	ValuesURL string                    `json:"values_url,omitempty"`
	Validate  []QueryParameterValidator `json:"validate,omitempty"`
	Format    string                    `json:"format,omitempty"`
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

type Function struct {
	Name       string              `json:"name"`
	Ident      string              `json:"ident,omitempty"`
	Required   []string            `json:"required"`
	Optional   []string            `json:"optional"`
	EnumSubset map[string][]string `json:"enum_subset,omitempty"`
	Examples   []string            `json:"examples"`
}
