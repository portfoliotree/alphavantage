package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/portfoliotree/alphavantage/specification"
)

func main() {
	idsBuf, err := os.ReadFile(filepath.FromSlash("specification/identifiers.json"))
	if err != nil {
		panic(err)
	}
	var goIdentifiers map[string][]string
	if err := json.Unmarshal(idsBuf, &goIdentifiers); err != nil {
		panic(err)
	}

	qpBuf, err := os.ReadFile(filepath.FromSlash("specification/query_parameters.json"))
	if err != nil {
		panic(err)
	}
	var queryParams []specification.QueryParameter
	if err := json.Unmarshal(qpBuf, &queryParams); err != nil {
		panic(err)
	}

	filePaths, err := filepath.Glob(filepath.FromSlash("specification/functions/*.json"))
	if err != nil {
		panic(err)
	}

	functionFiles := make(map[string][]specification.Function)

	for _, filePath := range filePaths {
		buf, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		var functions []specification.Function
		err = json.Unmarshal(buf, &functions)
		if err != nil {
			panic(err)
		}
		functionFiles[filePath] = functions
	}

	const pkgName = "alphavantage"

	for filePath, functions := range functionFiles {
		baseFileName := strings.TrimSuffix(filepath.Base(filePath), ".json")

		slices.SortFunc(functions, func(a, b specification.Function) int {
			return cmp.Compare(goIdentifiers[a.Name][0], goIdentifiers[b.Name][0])
		})

		outFileName := filepath.Join(baseFileName + ".go")

		if err := generateFile(pkgName, outFileName, baseFileName, functions, goIdentifiers, queryParams); err != nil {
			panic(err)
		}
	}
}

func generateFile(pkgName, outFileName, baseFileName string, functions []specification.Function, goIdentifiers map[string][]string, queryParams []specification.QueryParameter) error {
	file := ast.File{
		Name: ast.NewIdent(pkgName),
	}
	importsDecl := &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{},
	}
	file.Decls = append(file.Decls, importsDecl)

	imports := []string{
		"net/url",
	}

	for _, fn := range functions {
		goIdent := goIdentifiers[fn.Name][0]
		rowTypeIdent := goIdent + "Row"
		queryTypeIdent := goIdent + "Query"

		file.Decls = append(file.Decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(queryTypeIdent),
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("url"),
						Sel: ast.NewIdent("Values"),
					},
				},
			},
		})
		queryFuncDecls, im, err := queryInitializerFunc(goIdent, queryTypeIdent, fn, goIdentifiers, queryParams)
		if err != nil {
			return err
		}

		imports = append(imports, im...)
		slices.Sort(imports)
		imports = slices.Compact(imports)

		for _, fnDecl := range queryFuncDecls {
			file.Decls = append(file.Decls, fnDecl)
		}

		file.Decls = append(file.Decls, &ast.FuncDecl{
			Name: ast.NewIdent("Get" + goIdent),
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{ast.NewIdent("client")}, Type: &ast.StarExpr{X: ast.NewIdent("Client")}},
				},
			},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{Names: []*ast.Ident{ast.NewIdent("ctx")}, Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						}},
						{Names: []*ast.Ident{ast.NewIdent("q")}, Type: ast.NewIdent(queryTypeIdent)},
					},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("http"),
								Sel: ast.NewIdent("Response"),
							},
						}},
						{Type: ast.NewIdent("error")},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("req"),
							ast.NewIdent("err"),
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("client"),
									Sel: ast.NewIdent("newRequest"),
								},
								Args: []ast.Expr{
									ast.NewIdent("ctx"),
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   ast.NewIdent("url"),
											Sel: ast.NewIdent("Values"),
										},
										Args: []ast.Expr{ast.NewIdent("q")},
									},
								},
							},
						},
					},
					&ast.IfStmt{
						Cond: &ast.BinaryExpr{
							X:  ast.NewIdent("err"),
							Op: token.NEQ,
							Y:  ast.NewIdent("nil"),
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										ast.NewIdent("nil"),
										ast.NewIdent("err"),
									},
								},
							},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("res"),
							ast.NewIdent("err"),
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("client"),
									Sel: ast.NewIdent("Do"),
								},
								Args: []ast.Expr{
									ast.NewIdent("req"),
								},
							},
						},
					},
					&ast.IfStmt{
						Cond: &ast.BinaryExpr{
							X:  ast.NewIdent("err"),
							Op: token.NEQ,
							Y:  ast.NewIdent("nil"),
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										ast.NewIdent("nil"),
										ast.NewIdent("err"),
									},
								},
							},
						},
					},
					&ast.ReturnStmt{
						Results: []ast.Expr{
							ast.NewIdent("res"),
							ast.NewIdent("nil"),
						},
					},
				},
			},
		})

		imports = append(imports, "context", "net/http")
		slices.Sort(imports)
		imports = slices.Compact(imports)

		if !fn.HasDatatypeParameter() ||
			slices.Contains([]string{"REALTIME_BULK_QUOTES", "REALTIME_OPTIONS", "BOP"}, fn.Name) {
			continue
		}

		file.Decls = append(file.Decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(rowTypeIdent),
					Type: &ast.StructType{
						Fields: csvFields(baseFileName, fn, goIdentifiers),
					},
				},
			},
		}, &ast.FuncDecl{
			Name: ast.NewIdent("Get" + goIdent + "CSVRows"),
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{ast.NewIdent("client")}, Type: &ast.StarExpr{X: ast.NewIdent("Client")}},
				},
			},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{Names: []*ast.Ident{ast.NewIdent("ctx")}, Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						}},
						{Names: []*ast.Ident{ast.NewIdent("q")}, Type: ast.NewIdent(queryTypeIdent)},
					},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{Type: &ast.ArrayType{Elt: ast.NewIdent(rowTypeIdent)}},
						{Type: ast.NewIdent("error")},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("q"),
								Sel: ast.NewIdent("DataTypeCSV"),
							},
							Args: []ast.Expr{},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("res"),
							ast.NewIdent("err"),
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("client"),
									Sel: ast.NewIdent("Get" + goIdent),
								},
								Args: []ast.Expr{
									ast.NewIdent("ctx"),
									ast.NewIdent("q"),
								},
							},
						},
					},
					&ast.IfStmt{
						Cond: &ast.BinaryExpr{
							X:  ast.NewIdent("err"),
							Op: token.NEQ,
							Y:  ast.NewIdent("nil"),
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										ast.NewIdent("nil"),
										ast.NewIdent("err"),
									},
								},
							},
						},
					},
					&ast.DeferStmt{
						Call: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X:   ast.NewIdent("res"),
									Sel: ast.NewIdent("Body"),
								},
								Sel: ast.NewIdent("Close"),
							},
							Args: []ast.Expr{},
						},
					},
					&ast.DeclStmt{
						Decl: &ast.GenDecl{
							Tok: token.VAR,
							Specs: []ast.Spec{
								&ast.ValueSpec{
									Names: []*ast.Ident{ast.NewIdent("rows")},
									Type: &ast.ArrayType{
										Elt: ast.NewIdent(rowTypeIdent),
									},
								},
							},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("err"),
						},
						Tok: token.ASSIGN,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: ast.NewIdent("ParseCSV"),
								Args: []ast.Expr{
									&ast.SelectorExpr{
										X:   ast.NewIdent("res"),
										Sel: ast.NewIdent("Body"),
									},
									&ast.UnaryExpr{
										Op: token.AND,
										X:  ast.NewIdent("rows"),
									},
									ast.NewIdent("nil"),
								},
							},
						},
					},
					&ast.IfStmt{
						Cond: &ast.BinaryExpr{
							X:  ast.NewIdent("err"),
							Op: token.NEQ,
							Y:  ast.NewIdent("nil"),
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										ast.NewIdent("nil"),
										ast.NewIdent("err"),
									},
								},
							},
						},
					},
					&ast.ReturnStmt{
						Results: []ast.Expr{
							ast.NewIdent("rows"),
							ast.NewIdent("nil"),
						},
					},
				},
			},
		})
	}

	if len(file.Decls) <= 1 {
		return nil
	}

	slices.Sort(imports)
	imports = slices.Compact(imports)
	for _, im := range imports {
		importsDecl.Specs = append(importsDecl.Specs, &ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(im)}})
	}

	var buf bytes.Buffer
	buf.WriteString("// Code generated by github.com/portfoliotree/alphavantage/cmd/generate; DO NOT EDIT.\n")
	if err := format.Node(&buf, token.NewFileSet(), &file); err != nil {
		return err
	}

	return os.WriteFile(outFileName, buf.Bytes(), 0644)
}

func queryInitializerFunc(goIdent, queryTypeIdent string, fn specification.Function, goIdentifiers map[string][]string, queryParams []specification.QueryParameter) ([]*ast.FuncDecl, []string, error) {
	requiredFields := &ast.Field{
		Type: ast.NewIdent("string"),
	}
	if slices.Contains(fn.Required, "apikey") {
		requiredFields.Names = append(requiredFields.Names, ast.NewIdent(goIdentifiers["apikey"][1]))
	}

	elements := []ast.Expr{
		&ast.KeyValueExpr{
			Key: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote("function")},
			Value: &ast.CompositeLit{
				Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
				Elts: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(fn.Name)}},
			},
		},
	}

	for _, req := range fn.Required {
		switch req {
		case "function":
		default:
			requiredFields.Names = append(requiredFields.Names, ast.NewIdent(goIdentifiers[req][1]))
			fallthrough
		case "apikey":
			elements = append(elements, &ast.KeyValueExpr{
				Key: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(req)},
				Value: &ast.CompositeLit{
					Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
					Elts: []ast.Expr{ast.NewIdent(goIdentifiers[req][1])},
				},
			})
		}
	}

	decls := []*ast.FuncDecl{
		{
			Name: ast.NewIdent("Query" + goIdent),
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						requiredFields,
					},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{Type: ast.NewIdent(queryTypeIdent)},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.CompositeLit{
								Type: ast.NewIdent(queryTypeIdent),
								Elts: elements,
							},
						},
					},
				},
			},
		},
	}

	var imports []string

	for _, opt := range fn.Optional {
		var methodName string
		switch opt {
		case "OHLC":
			methodName = "OHLC"
		default:
			ids, ok := goIdentifiers[opt]
			if ok {
				methodName = ids[0]
			} else {
				panic("unknown option: " + opt)
			}
		}

		if idx := slices.IndexFunc(queryParams, func(p specification.QueryParameter) bool {
			return p.Name == opt
		}); idx >= 0 {

			qp := queryParams[idx]
			switch qp.Type {
			case "enum":
				for _, val := range qp.Values {
					s, err := val.Name()
					if err != nil {
						return nil, nil, fmt.Errorf("failed to get name for query param key %s: %w", opt, err)
					}
					valueIdent, ok := goIdentifiers[s]
					if ok {
						decls = append(decls, queryUpdateMethod(queryTypeIdent, methodName+valueIdent[0], opt,
							[]*ast.Field{},
							&ast.CompositeLit{
								Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
								Elts: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(s)}},
							},
						))
					}
				}
			case "bool":
				decls = append(decls, queryUpdateMethod(queryTypeIdent, methodName, opt,
					[]*ast.Field{
						{
							Names: []*ast.Ident{ast.NewIdent("value")},
							Type:  ast.NewIdent("bool"),
						},
					},
					&ast.CompositeLit{
						Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
						Elts: []ast.Expr{&ast.CallExpr{
							Fun:  &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("FormatBool")},
							Args: []ast.Expr{ast.NewIdent("value")},
						}},
					},
				))
				imports = append(imports, "strconv")
				continue
			case "time":
				if qp.Format != "" {
					decls = append(decls, queryUpdateMethod(queryTypeIdent, methodName, opt,
						[]*ast.Field{
							{
								Names: []*ast.Ident{ast.NewIdent("value")},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Time"),
								},
							},
						},
						&ast.CompositeLit{
							Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
							Elts: []ast.Expr{&ast.CallExpr{
								Fun:  &ast.SelectorExpr{X: ast.NewIdent("value"), Sel: ast.NewIdent("Format")},
								Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(qp.Format)}},
							}},
						},
					))
					imports = append(imports, "time")
				}
				continue
			}
		}
		decls = append(decls, queryUpdateMethod(queryTypeIdent, methodName, opt,
			[]*ast.Field{{
				Names: []*ast.Ident{ast.NewIdent("value")},
				Type:  ast.NewIdent("string"),
			}},
			&ast.CompositeLit{
				Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
				Elts: []ast.Expr{ast.NewIdent("value")},
			},
		))
	}

	slices.Sort(imports)
	imports = slices.Compact(imports)

	return decls, imports, nil
}

func queryUpdateMethod(queryTypeIdent, methodName, keyName string, params []*ast.Field, expr ast.Expr) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(methodName),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent("query")}, Type: ast.NewIdent(queryTypeIdent)},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: params,
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{Type: ast.NewIdent(queryTypeIdent)}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.IndexExpr{
							X:     ast.NewIdent("query"),
							Index: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(keyName)},
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{expr},
				},
				&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent("query")}},
			},
		},
	}
}

func csvFields(baseFileName string, fn specification.Function, goIdentifiers map[string][]string) *ast.FieldList {
	if strings.HasPrefix(baseFileName, "technical_") && len(fn.CSVColumns) == 2 {
		return &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent(goIdentifiers[fn.CSVColumns[0]][0])},
					Type:  ast.NewIdent("string"),
					Tag: &ast.BasicLit{
						Kind:  token.STRING,
						Value: "`" + fmt.Sprintf(`column-name:%q`, fn.CSVColumns[0]) + "`",
					},
				},
				{
					Names: []*ast.Ident{ast.NewIdent("Value")},
					Type:  ast.NewIdent("string"),
					Tag: &ast.BasicLit{
						Kind:  token.STRING,
						Value: "`" + fmt.Sprintf(`column-name:%q`, fn.CSVColumns[1]) + "`",
					},
				},
			},
		}
	}

	switch fn.Name {
	default:
		fields := ast.FieldList{}
		for _, col := range fn.CSVColumns {
			fields.List = append(fields.List, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(goIdentifiers[col][0])},
				Type:  ast.NewIdent("string"),
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`" + fmt.Sprintf(`column-name:%q`, col) + "`",
				},
			})
		}
		return &fields
	}
}
