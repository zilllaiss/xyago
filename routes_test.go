package main

import (
	"testing"
)

func TestUtils(t *testing.T) {
	md, err := parseMarkdownFiles("posts")
	if err != nil || len(md) < 1 {
		t.Fatalf("err: %v, len: %d", err.Error(), len(md))
	}
	for _, v := range md {
		t.Logf("v: %#v\n", v.Frontmatter)
	}

	if err := sortMarkdownFiles(md); err != nil {
		t.Fatalf("err: %v", err.Error())
	}
	for _, v := range md {
		t.Logf("title: %v\ndate: %v", v.Slug, v.Frontmatter.PublishedAt)
	}
}
