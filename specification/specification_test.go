package specification_test

import (
	"encoding/csv"
	"encoding/json"
	"go/token"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage/specification"
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
			formatted, err := json.MarshalIndent(json.RawMessage(buf), "", specification.JSONIndent)
			require.NoError(t, err)
			require.NoError(t, os.WriteFile(match, formatted, info.Mode().Perm()))
		}
	}
}

func TestQueryParameters(t *testing.T) {
	buf, err := os.ReadFile("query_parameters.json")
	require.NoError(t, err)
	var data []specification.QueryParameter
	require.NoError(t, json.Unmarshal(buf, &data))

	for _, param := range data {
		t.Run(param.Name, func(t *testing.T) {
			require.NotEmpty(t, param.Type)
			require.NotEmpty(t, param.Name)
			validateQueryParameterType(t, param)
		})
	}
}

func validateQueryParameterType(t *testing.T, param specification.QueryParameter) {
	t.Helper()
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

func TestIdentifiers(t *testing.T) {
	buf, err := os.ReadFile("identifiers.json")
	require.NoError(t, err)
	var data map[string][]string
	require.NoError(t, json.Unmarshal(buf, &data))

	for avID, goIDs := range data {
		t.Run(avID, func(t *testing.T) {
			require.Truef(t, token.IsIdentifier(goIDs[0]), "expected %s to be an identifier", goIDs[0])
			switch goIDs[1] {
			case "rng", "tp": // exceptions
			default:
				require.Equal(t, strings.ToLower(goIDs[0]), strings.ToLower(goIDs[1]))
			}
			require.Truef(t, token.IsIdentifier(goIDs[1]), "expected %s to be an identifier", goIDs[1])
			require.Truef(t, token.IsExported(goIDs[0]), "expected %s to be an public", goIDs[0])
			require.Falsef(t, token.IsExported(goIDs[1]), "expected %s to be an private", goIDs[1])
			require.NotContains(t, goIDs[0], "_")
			require.NotContains(t, goIDs[0], " ")
			require.NotContains(t, goIDs[1], "_")
			require.NotContains(t, goIDs[1], " ")
		})
	}
}

func TestFunctions(t *testing.T) {
	qpBuf, err := os.ReadFile("query_parameters.json")
	require.NoError(t, err, "failed to read query_parameters.json")
	var queryParameters []specification.QueryParameter
	require.NoError(t, json.Unmarshal(qpBuf, &queryParameters), "failed to parse query_parameters.json")

	idsBuf, err := os.ReadFile("identifiers.json")
	require.NoError(t, err)
	var goIdentifiers map[string][]string
	require.NoError(t, json.Unmarshal(idsBuf, &goIdentifiers))

	functionFiles := loadFunctions(t)
	for filePath, functions := range functionFiles {
		baseFileName := strings.TrimSuffix(filepath.Base(filePath), ".json")
		t.Run(baseFileName, func(t *testing.T) {
			t.Parallel()
			for _, fn := range functions {
				t.Run(fn.Name, func(t *testing.T) {
					require.NotEmpty(t, fn.Name)
					allQueryParamsSpecified(t, fn, queryParameters, fn.Required...)
					allQueryParamsSpecified(t, fn, queryParameters, fn.Optional...)
					enumValuesSubset(t, fn, queryParameters)
					testExampleURLs(t, fn)

					if _, found := goIdentifiers[fn.Name]; assert.Truef(t, found, "go identifiers must have key for %q", fn.Name) {
					}

					for _, val := range fn.Required {
						if _, found := goIdentifiers[fn.Name]; assert.Truef(t, found, "go identifiers must have key for required query param key %q", val) {
						}
					}
					for _, val := range fn.Optional {
						if _, found := goIdentifiers[fn.Name]; assert.Truef(t, found, "go identifiers must have key for optional query param key %q", val) {
						}
					}
				})
			}
		})
	}
}

func TestCSVColumns(t *testing.T) {
	apikey, hasAPIKey := os.LookupEnv("ALPHA_VANTAGE_TOKEN")
	if !hasAPIKey || apikey == "" {
		t.Skip("skipping test because env var ALPHA_VANTAGE_TOKEN is not set")
		return
	}

	exampleDir := filepath.FromSlash("testdata/examples")

	indexEntrees := loadTestdataExampleIndex(t)
	indexEntrees = removeMissingExampleFileEntries(t, indexEntrees)

	for filePath, functions := range loadFunctions(t) {
		baseFileName := strings.TrimSuffix(filepath.Base(filePath), ".json")
		require.NoError(t, os.MkdirAll(filepath.Join(exampleDir, baseFileName), 0744))

		columnsSet := false

		t.Run(baseFileName, func(t *testing.T) {
			for fi, fn := range functions {
				for _, exampleURL := range fn.Examples {
					indexEntrees = testdataExampleBody(t, filepath.Join(exampleDir, baseFileName), apikey, indexEntrees, fn, exampleURL, func(bodyFilepath string) {
						if filepath.Ext(bodyFilepath) != ".csv" {
							return
						}

						f, err := os.Open(bodyFilepath)
						require.NoError(t, err)
						t.Cleanup(func() {
							_ = f.Close()
						})
						r := csv.NewReader(f)

						firstLine, err := r.Read()
						require.NoError(t, err)

						var columnNames []string
						for _, n := range fn.CSVColumns {
							columnNames = append(columnNames, n.Name)
						}

						if !assert.Equal(t, columnNames, firstLine) && len(fn.CSVColumns) == 0 {
							for _, name := range firstLine {
								functions[fi].CSVColumns = append(functions[fi].CSVColumns, specification.CSVColumn{Name: name, Type: "string"})
							}
							columnsSet = true
						}
					})
				}
			}
		})

		if columnsSet {
			buf, err := json.MarshalIndent(functions, "", specification.JSONIndent)
			require.NoError(t, err)
			require.NoError(t, os.WriteFile(filePath, buf, 0644))
		}
	}

	saveTestdataExampleIndex(t, indexEntrees)
}

func allQueryParamsSpecified(t *testing.T, fn specification.Function, queryParameters []specification.QueryParameter, names ...string) {
	t.Helper()
	for _, param := range names {
		index := slices.IndexFunc(queryParameters, func(parameter specification.QueryParameter) bool {
			return parameter.Name == param
		})
		if index < 0 {
			t.Errorf("failed to find required parameter %q for %s", param, fn.Name)
		}
	}
}

func enumValuesSubset(t *testing.T, fn specification.Function, queryParameters []specification.QueryParameter) {
	if len(fn.EnumSubset) == 0 {
		return
	}
	for key, values := range fn.EnumSubset {
		index := slices.IndexFunc(queryParameters, func(parameter specification.QueryParameter) bool {
			return parameter.Name == key
		})
		if index < 0 {
			t.Errorf("failed to find required parameter %q for %s", key, fn.Name)
		}
		paramSpec := queryParameters[index]

		var specValues []string

		for _, value := range paramSpec.Values {
			s, err := value.Name()
			require.NoError(t, err)
			specValues = append(specValues, s)
		}

		for _, value := range values {
			i := slices.Index(specValues, value)
			if i < 0 {
				t.Errorf("failed to find required parameter %q in %q", value, specValues)
			}
		}
	}
}

func testExampleURLs(t *testing.T, fn specification.Function) {
	t.Helper()
	for _, example := range fn.Examples {
		u, err := url.Parse(example)
		require.NoError(t, err)

		q := u.Query()

		require.Equal(t, fn.Name, q.Get("function"))
	}
}

func loadFunctions(t *testing.T) map[string][]specification.Function {
	t.Helper()

	filePaths, err := filepath.Glob(filepath.FromSlash("functions/*.json"))
	require.NoError(t, err)

	files := make(map[string][]specification.Function)

	for _, filePath := range filePaths {
		buf, err := os.ReadFile(filePath)
		require.NoError(t, err)
		var functions []specification.Function
		require.NoError(t, json.Unmarshal(buf, &functions))

		files[filePath] = functions
	}

	return files
}
