package main

import (
	"cmp"
	"fmt"
	"slices"
	"time"

	"xyago/types"

	"codeberg.org/zill_laiss/fest/markdown"
	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	tagsMap      = map[string][]*markdown.MarkdownData{}
	tagsSorted   = []string{}
	posts        []*markdown.MarkdownData
	dynamicPosts []*markdown.MarkdownData
)

func mainMdParser() *WrappedMarkdownParser {
	m := markdown.NewMarkdown(
		goldmark.WithRendererOptions(html.WithUnsafe()),
		goldmark.WithExtensions(
			figure.Figure,
			extension.Table,
			extension.DefinitionList,
		),
	)
	return &WrappedMarkdownParser{mp: m}
}

type WrappedMarkdownParser struct {
	mp markdown.MarkdownParser
}

func (wmp *WrappedMarkdownParser) ParseFile(path string) (*markdown.MarkdownData, error) {
	f, err := wmp.mp.ParseFile(path)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %w", err)
	}

	fm := &types.Frontmatter{}
	if err := f.GetFrontmatter(fm); err != nil {
		return nil, fmt.Errorf("error getting frontmatter: %w", err)
	}

	if !fm.Public {
		f = nil
	} else {
		for _, t := range fm.Tags {
			if tagsMap[t] == nil {
				tagsMap[t] = []*markdown.MarkdownData{}
			}
			tagsMap[t] = append(tagsMap[t], f)
		}
	}
	return f, nil
}

func (wmp *WrappedMarkdownParser) ParseFiles(path string) ([]*markdown.MarkdownData, error) {
	m := []*markdown.MarkdownData{}
	if err := markdown.ScanForMarkdown(path, func(p string) error {
		f, err := wmp.ParseFile(p)
		if err != nil {
			return fmt.Errorf("error parsing file: %w", err)
		}
		if f != nil {
			m = append(m, f)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("error scanning markdown, %w", err)
	}
	return m, nil
}

func sortMarkdownFilesByDate(md []*markdown.MarkdownData, considerUpdatedDate bool) error {
	var sortErr error
	slices.SortStableFunc(md, func(a, b *markdown.MarkdownData) int {
		aFm := &types.Frontmatter{}
		bFm := &types.Frontmatter{}

		err1 := a.GetFrontmatter(aFm)
		err2 := b.GetFrontmatter(bFm)

		if err := cmp.Or(err1, err2); err != nil {
			sortErr = err
			return 0
		}

		var parseError error

		updatedOrPublished := func(fm *types.Frontmatter) *time.Time {
			if parseError != nil {
				return nil
			}

			var t time.Time
			var err error

			if considerUpdatedDate {
				t, err = time.Parse(time.DateOnly, fm.UpdatedAt)
				if err == nil {
					return &t
				}
			}

			t, err = time.Parse(time.DateOnly, fm.PublishedAt)
			if err == nil {
				return &t
			}
			parseError = err
			return nil
		}

		t1 := updatedOrPublished(aFm)
		t2 := updatedOrPublished(bFm)
		if parseError != nil {
			sortErr = parseError
			return 0
		}

		return int((*t1).Sub(*t2)) * -1
	})
	return sortErr
}
