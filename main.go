package main

import (
	"log"
)

func main() {
	var err error

	posts, err = parseMarkdownFiles("posts")
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
