package main

import (
	"log"

	md "github.com/zilllaiss/fest/markdown"
)

func main() {
	var err error

	mdp := md.NewMarkdown()

	posts, err = mdp.ParseFiles("posts")
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
