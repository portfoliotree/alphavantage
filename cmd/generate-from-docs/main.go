package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"maps"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/crhntr/dom"
	"github.com/crhntr/dom/spec"
	"golang.org/x/net/html"
)

func main() {
	document, err := parseDocumentationPage(context.Background(), time.Hour)
	if err != nil {
		log.Fatal(err)
	}
	for section := range document.QuerySelectorSequence("section") {
		h2 := section.QuerySelector(`h2`)
		if h2 == nil {
			continue
		}
		fmt.Printf("%s https://www.alphavantage.co/documentation/#%s\n", strings.TrimSpace(h2.TextContent()), h2.ID())

		fns := make(map[string]Function)

		for a := range section.QuerySelectorSequence(`a[href^="https://www.alphavantage.co/query"]`) {
			href, err := url.Parse(a.GetAttribute("href"))
			if err != nil {
				log.Fatal(err)
			}
			function := href.Query().Get("function")
			if function == "" {
				continue
			}
			fn, ok := fns[function]
			if ok {
				continue
			}
			for k := range href.Query() {
				fn.Param = append(fn.Param, k)
			}
			slices.Sort(fn.Param)
			fn.Param = slices.Compact(fn.Param)
			for prev := a.ParentElement().PreviousSibling(); prev != nil; prev = prev.PreviousSibling() {
				el, ok := prev.(spec.Element)
				if !ok || !el.Matches(`h4`) {
					continue
				}
				fn.Name = strings.TrimSpace(el.TextContent())
				fn.URL = "https://www.alphavantage.co/documentation/#" + el.ID()
			}
			fns[function] = fn
		}

		fnNames := slices.Collect(maps.Keys(fns))
		slices.Sort(fnNames)
		fnNames = slices.Compact(fnNames)

		for _, fnName := range fnNames {
			fn := fns[fnName]
			fmt.Printf("  [%s] %s %s\n", fnName, fn.Name, fn.URL)
			fmt.Printf("    %s\n", strings.Join(fn.Param, ", "))
		}
	}
}

type Function struct {
	Name  string
	URL   string
	Param []string
}

func parseDocumentationPage(ctx context.Context, ttl time.Duration) (spec.Document, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user cache dir: %v", err)
	}
	docsBuf, err := readDocumentationPage(ctx, filepath.Join(cacheDir, filepath.FromSlash("github.com/portfoliotree/alphavantage"), "documentation.html"), ttl)
	if err != nil {
		return nil, err
	}
	defer closeAndIgnoreError(docsBuf)
	documentNode, err := html.Parse(docsBuf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}
	if n := dom.NewNode(documentNode); n == nil {
		return nil, fmt.Errorf("failed to create document node")
	} else if document, ok := n.(spec.Document); !ok {
		return nil, fmt.Errorf("failed to create document node")
	} else {
		return document, nil
	}
}

func readDocumentationPage(ctx context.Context, cacheFile string, ttl time.Duration) (io.ReadCloser, error) {
	dir := filepath.Dir(cacheFile)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}
	if info, err := os.Stat(cacheFile); err != nil || time.Since(info.ModTime()) > ttl {
		file, err := os.Create(cacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to create cache file: %w", err)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.alphavantage.co/documentation/", nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		if res, err := http.DefaultClient.Do(req); err != nil {
			log.Fatal(err)
		} else {
			if res.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("unexpected status code: %s (%d)", http.StatusText(res.StatusCode), res.StatusCode)
			}
			defer closeAndIgnoreError(res.Body)
			if _, err := io.Copy(file, res.Body); err != nil {
				return nil, fmt.Errorf("failed to copy response from server: %w", err)
			}
		}
		if _, seekErr := file.Seek(0, io.SeekStart); seekErr != nil {
			return nil, fmt.Errorf("failed to seek in cache file: %w", seekErr)
		}
		return file, nil
	}
	return os.Open(cacheFile)
}

func closeAndIgnoreError(c io.Closer) {
	if err := c.Close(); err != nil {
		slog.Debug("close error", slog.String("message", err.Error()))
	}
}
