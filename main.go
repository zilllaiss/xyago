package main

import (
	"log"

	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/zilllaiss/fest/markdown"
)

func main() {
	var err error

	m := markdown.NewMarkdown(
		goldmark.WithRendererOptions(html.WithUnsafe()),
		goldmark.WithExtensions(figure.Figure),
	)
	wmp := &WrappedMarkdownParser{mp: m}

	posts, err = wmp.ParseFiles("posts")
	if err != nil {
		log.Fatalln(err)
	}

	if err := sortMarkdownFiles(posts); err != nil {
		log.Fatalln(err)
	}

	g := generator()

	if err := g.Generate(); err != nil {
		log.Fatalln(err)
	}
}
