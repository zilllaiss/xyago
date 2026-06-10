package main

import (
	"testing"

	"xyago/types"
)

func TestUtils(t *testing.T) {
	wmp := mainMdParser()
	md, err := wmp.ParseFiles("posts")
	if err != nil || len(md) < 1 {
		t.Fatalf("err: %v, len: %d", err.Error(), len(md))
	}
	for _, v := range md {
		var fm types.Frontmatter
		if err := v.GetFrontmatter(&fm); err != nil {
			t.Fatalf("%v", err.Error())
		}
		t.Logf("v: %#v\n", fm)
	}

	if err := sortMarkdownFilesByDate(md, false); err != nil {
		t.Fatalf("err: %v", err.Error())
	}
	for _, v := range md {
		var fm types.Frontmatter
		if err := v.GetFrontmatter(&fm); err != nil {
			t.Fatalf("%v", err.Error())
		}
		t.Logf("title: %v\ndate: %v", v.Slug, fm.PublishedAt)
	}
}
