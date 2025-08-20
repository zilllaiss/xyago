package main

import (
	"cmp"
	"fmt"
	"slices"
	"time"
	"xyago/types"

	"github.com/zilllaiss/fest/markdown"
)

var (
	tagsMap    = map[string][]*markdown.MarkdownData{}
	tagsSorted = []string{}
	posts      []*markdown.MarkdownData
)

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
		m = append(m, f)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("error scanning markdown, %w", err)
	}

	for k := range tagsMap {
		tagsSorted = append(tagsSorted, k)
	}

	slices.Sort(tagsSorted)
	return m, nil
}

func sortMarkdownFiles(md []*markdown.MarkdownData) error {
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

		t1, err1 := time.Parse(time.DateOnly, aFm.PublishedAt)
		t2, err2 := time.Parse(time.DateOnly, bFm.PublishedAt)

		if err := cmp.Or(err1, err2); err != nil {
			sortErr = err
			return 0
		}

		return int(t1.Sub(t2)) * -1
	})
	return sortErr
}
