package main

import (
	"bytes"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"xyago/types"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/frontmatter"
	"go.abhg.dev/goldmark/toc"
)

var (
	markdown   = goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{}))
	tagsMap    = map[string][]*types.Markdown{}
	tagsSorted = []string{}
	posts      []types.Markdown
)

func parseMarkdownFiles(path string) ([]types.Markdown, error) {
	mdDir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	ctx := parser.NewContext()
	markdown.Parser().AddOptions(parser.WithAutoHeadingID())

	fm := []types.Markdown{}

	for _, content := range mdDir {
		if ext := filepath.Ext(content.Name()); content.IsDir() || ext != ".md" {
			continue
		}
		filename := filepath.Join(path, content.Name())

		md, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var mdBuf bytes.Buffer
		var tocBuf bytes.Buffer

		doc := markdown.Parser().Parse(text.NewReader(md), parser.WithContext(ctx))

		tree, err := toc.Inspect(doc, md, toc.MaxDepth(3), toc.Compact(true))
		if err != nil {
			return nil, err
		}

		list := toc.RenderList(tree)
		if list != nil {
			if err := markdown.Renderer().Render(&tocBuf, md, list); err != nil {
				return nil, err
			}
		}

		if err := markdown.Renderer().Render(&mdBuf, md, doc); err != nil {
			return nil, err
		}

		var festMd types.Markdown
		ptrFestMd := &festMd

		metaData := frontmatter.Get(ctx)
		if err := metaData.Decode(&ptrFestMd.Frontmatter); err != nil {
			return nil, err
		}

		for _, t := range festMd.Frontmatter.Tags {
			if tagsMap[t] == nil {
				tagsMap[t] = []*types.Markdown{}
			}
			tagsMap[t] = append(tagsMap[t], &festMd)
		}
		ptrFestMd.Slug = strings.TrimSuffix(content.Name(), ".md")
		ptrFestMd.Content = mdBuf.String()
		ptrFestMd.TOC = tocBuf.String()

		fm = append(fm, festMd)
	}

	for k := range tagsMap {
		tagsSorted = append(tagsSorted, k)
	}
	slices.Sort(tagsSorted)

	return fm, nil
}

func sortMarkdownFiles(md []types.Markdown) error {
	var sortErr error
	slices.SortStableFunc(md, func(a, b types.Markdown) int {
		t1, err := time.Parse(time.DateOnly, a.Frontmatter.PublishedAt)
		if err != nil {
			sortErr = err
			return 0
		}
		t2, err := time.Parse(time.DateOnly, b.Frontmatter.PublishedAt)
		if err != nil {
			sortErr = err
			return 0
		}

		return int(t1.Sub(t2)) * -1
	})
	return sortErr
}
