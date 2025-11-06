// scan-docs should be run from the repository root
package main

import (
	"bytes"
	_ "embed"
	"io"
	"log/slog"
	"maps"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/typelate/dom"
	"github.com/typelate/dom/spec"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

type Function struct {
	Name     string
	Category string
	Required []string
	Optional []string
	Values   url.Values
	Examples []string
}

const (
	htmlFile = "internal/scan-docs/documentation.html"
	yamlFile = "internal/scan-docs/documentation.yaml"
)

func main() {
	var buf []byte
	if info, err := os.Stat(filepath.FromSlash(htmlFile)); err != nil || info.ModTime().Before(time.Now().AddDate(0, 0, 7)) {
		documentationResponse, err := http.Get("https://www.alphavantage.co/documentation/")
		if err != nil {
			panic(err)
		}
		defer closeAndIgnoreError(documentationResponse.Body)
		buf, err = io.ReadAll(documentationResponse.Body)
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile(filepath.FromSlash(htmlFile), buf, 0644); err != nil {
			panic(err)
		}
	} else {
		buf, err = os.ReadFile(filepath.FromSlash(htmlFile))
		if err != nil {
			panic(err)
		}
	}

	dn, err := html.Parse(bytes.NewBuffer(buf))
	if err != nil {
		panic(err)
	}
	document := dom.NewNode(dn).(spec.Document)

	var allFunctionNames []string

	var functions []Function

	for endpointCategory := range document.QuerySelectorSequence("main article section h2") {
		elemList := endpointCategory.Closest("section").Children()
		categoryName := endpointCategory.TextContent()

		for _, f1 := range splitByMatch(elemList, currySecondParamOneResult(spec.Element.Matches, `h4`)) {
			for _, apiParams := range splitByMatch(f1.Children(), currySecondParamOneResult(spec.Element.Matches, `h6`)) {
				h6 := apiParams.QuerySelector(`h6`)
				if h6 == nil {
					continue
				}
				if strings.TrimSpace(h6.TextContent()) != "API Parameters" {
					continue
				}
				var (
					required, optional []string
					values             = make(url.Values)
				)

				const paramKeyPrefix = "❚ "
				children := apiParams.Children()
				for i := 0; i < children.Length(); i++ {
					h := children.Item(i)
					content := strings.TrimSpace(h.TextContent())
					if !strings.HasPrefix(content, paramKeyPrefix) {
						continue
					}
					if _, after, ok := strings.Cut(content, ": "); ok {
						key := strings.TrimSpace(after)
						if strings.HasPrefix(content, "❚ Required: ") {
							required = append(required, key)
						} else if strings.HasPrefix(content, "❚ Optional: ") {
							optional = append(optional, key)
						}
						values[key] = []string{}
					}
				}

				for codeEl := range apiParams.QuerySelectorSequence(`code`) {
					codeContent := strings.TrimSpace(codeEl.TextContent())
					if strings.ContainsAny(codeContent, "\n():") {
						continue
					}
					before, after, ok := strings.Cut(codeContent, "=")
					if !ok {
						continue
					}
					values[before] = append(values[before], strings.TrimSpace(after))
					if p := codeEl.Closest("p"); p != nil {
						for example := range p.QuerySelectorSequence(`code`) {
							content := strings.TrimSpace(example.TextContent())
							if strings.Contains(content, "=") || content == before {
								continue
							}
							values[before] = append(values[before], content)
						}
					}
				}

				funcName := values.Get("function")
				if funcName == "" {
					continue
				}
				keys := slices.Collect(maps.Keys(values))
				slices.Sort(keys)

				for _, key := range keys {
					s := values[key]
					slices.Sort(s)
					s = slices.Compact(s)
					values[key] = s
				}

				allFunctionNames = append(allFunctionNames, funcName)
				functions = append(functions, Function{
					Name:     funcName,
					Category: categoryName,
					Required: required,
					Optional: optional,
					Values:   values,
				})
			}
		}
	}

	for a := range document.QuerySelectorSequence(`a[href^="https://www.alphavantage.co/query"]`) {
		href := a.GetAttribute("href")
		u, err := url.Parse(href)
		if err != nil {
			continue
		}
		fn := u.Query().Get("function")
		idx := slices.IndexFunc(functions, func(function Function) bool {
			return function.Name == fn
		})
		if idx < 0 {
			continue
		}
		functions[idx].Examples = append(functions[idx].Examples, href)
	}

	slices.Sort(allFunctionNames)
	allFunctionNames = slices.Compact(allFunctionNames)

	outYAML, err := yaml.Marshal(functions)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.FromSlash(yamlFile), outYAML, 0644); err != nil {
		panic(err)
	}
}

func currySecondParamOneResult[P1, P2, R1 any](f func(P1, P2) R1, p2 P2) func(P1) R1 {
	return func(p1 P1) R1 { return f(p1, p2) }
}

func splitByMatch(collection spec.ElementCollection, match func(el spec.Element) bool) []spec.DocumentFragment {
	var fragments []spec.DocumentFragment
	maxChildIndex := collection.Length()
	start := -1
	for i := 0; i < maxChildIndex; i++ {
		if match(collection.Item(i)) {
			start = i
		}
	}
	if start < 0 {
		return nil
	}
	for i := 0; i < maxChildIndex; i++ {
		if match(collection.Item(i)) {
			fragments = append(fragments, elementsWithinRange(collection, start, i))
			start = i
		}
	}
	if start < maxChildIndex && match(collection.Item(start)) {
		fragments = append(fragments, elementsWithinRange(collection, start, maxChildIndex))
	}
	return fragments
}

func elementsWithinRange(col spec.ElementCollection, start, end int) spec.DocumentFragment {
	fragment := dom.NewDocumentFragment(nil)
	for j := start; j < end; j++ {
		fragment.Append(col.Item(j))
	}
	return fragment
}

func closeAndIgnoreError(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		slog.Debug("error closing body", slog.String("message", err.Error()))
	}
}
