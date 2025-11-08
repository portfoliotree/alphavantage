package specification

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestJSON just formats the files and makes sure they have valid JSON
func TestJSON(t *testing.T) {
	for _, pattern := range []string{
		"*.json",
		"functions/*.json",
	} {
		matches, err := filepath.Glob(pattern)
		require.NoError(t, err)
		for _, match := range matches {
			info, err := os.Stat(match)
			require.NoError(t, err)
			buf, err := os.ReadFile(match)
			require.NoError(t, err)
			formatted, err := json.MarshalIndent(json.RawMessage(buf), "", "\t")
			require.NoError(t, err)
			require.NoError(t, os.WriteFile(match, formatted, info.Mode().Perm()))
		}
	}
}

func TestInitialisms(t *testing.T) {
	buf, err := os.ReadFile("initialisms.json")
	require.NoError(t, err)
	var data []string
	require.NoError(t, json.Unmarshal(buf, &data))
	require.Truef(t, slices.IsSorted(data), "expected values to be sorted alphabetically")
}

func TestCompoundWords(t *testing.T) {
	buf, err := os.ReadFile("compound_words.json")
	require.NoError(t, err)
	var data map[string][]string
	require.NoError(t, json.Unmarshal(buf, &data))
}

func TestQueryParameters(t *testing.T) {
	buf, err := os.ReadFile("query_parameters.json")
	require.NoError(t, err)
	var data []QueryParameter
	require.NoError(t, json.Unmarshal(buf, &data))

	for _, param := range data {
		t.Run(param.Name, func(t *testing.T) {
			require.NotEmpty(t, param.Type)
			require.NotEmpty(t, param.Name)
			validateQueryParameterType(t, param)
		})
	}
}

func validateQueryParameterType(t *testing.T, param QueryParameter) {
	switch param.Type {
	case "bool":
		require.Empty(t, param.ValuesURL)
	case "comma_separated_enum":
		require.Empty(t, param.Format)
		require.NotEmpty(t, param.Values)
	case "comma_separated_list":
		require.Empty(t, param.Format)
	case "time":
		require.NotEmpty(t, param.Format)
	case "enum":
		require.Empty(t, param.Format)
		require.NotEmpty(t, param.Values)
	case "float":
		require.Empty(t, param.Format)
	case "int":
		require.Empty(t, param.Format)
	case "string":
		require.Empty(t, param.Format)
	default:
		t.Errorf("unknown query parameter type: %s", param.Type)
	}
}

func TestFunctions(t *testing.T) {
	buf, err := os.ReadFile("query_parameters.json")
	require.NoError(t, err, "failed to read query_parameters.json")
	var queryParameters []QueryParameter
	require.NoError(t, json.Unmarshal(buf, &queryParameters), "failed to parse query_parameters.json")

	filePaths, err := filepath.Glob(filepath.FromSlash("functions/*.json"))
	require.NoError(t, err)
	for _, filePath := range filePaths {
		baseFileName := strings.TrimSuffix(filepath.Base(filePath), ".json")
		t.Run(baseFileName, func(t *testing.T) {
			buf, err := os.ReadFile(filePath)
			require.NoError(t, err)
			var functions []Function
			require.NoError(t, json.Unmarshal(buf, &functions))

			for _, fn := range functions {
				t.Run(fn.Name, func(t *testing.T) {
					require.NotEmpty(t, fn.Name)
					allQueryParamsSpecified(t, fn, queryParameters, fn.Required...)
					allQueryParamsSpecified(t, fn, queryParameters, fn.Optional...)
					enumValuesSubset(t, fn, queryParameters)
					testExampleURLs(t, fn)
				})
			}
		})
	}
}

func allQueryParamsSpecified(t *testing.T, fn Function, queryParameters []QueryParameter, names ...string) {
	t.Helper()
	for _, param := range names {
		index := slices.IndexFunc(queryParameters, func(parameter QueryParameter) bool {
			return parameter.Name == param
		})
		if index < 0 {
			t.Errorf("failed to find required parameter %q for %s", param, fn.Name)
		}
	}
}

func enumValuesSubset(t *testing.T, fn Function, queryParameters []QueryParameter) {
	if len(fn.EnumSubset) == 0 {
		return
	}
	for key, values := range fn.EnumSubset {
		index := slices.IndexFunc(queryParameters, func(parameter QueryParameter) bool {
			return parameter.Name == key
		})
		if index < 0 {
			t.Errorf("failed to find required parameter %q for %s", key, fn.Name)
		}
		paramSpec := queryParameters[index]

		var specValues []string

		for _, value := range paramSpec.Values {
			switch value[0] {
			case '{':
				var elem struct {
					Value string `json:"value"`
				}
				require.NoError(t, json.Unmarshal([]byte(value), &elem))
				specValues = append(specValues, elem.Value)
			case '"':
				var str string
				require.NoError(t, json.Unmarshal([]byte(value), &str))
				specValues = append(specValues, str)
			default:
				t.Errorf("unknown query parameter value shape `%s` for %s", value, fn.Name)
			}
		}

		for _, value := range values {
			i := slices.Index(specValues, value)
			if i < 0 {
				t.Errorf("failed to find required parameter %q in %q", value, specValues)
			}
		}
	}
}

func testExampleURLs(t *testing.T, fn Function) {
	for _, example := range fn.Examples {
		u, err := url.Parse(example)
		require.NoError(t, err)

		q := u.Query()

		require.Equal(t, fn.Name, q.Get("function"))
	}
}
