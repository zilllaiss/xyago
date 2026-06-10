package main

import (
	"cmp"
	"log"
	"slices"
)

func main() {
	var err error

	wmp := mainMdParser()

	var err1, err2 error

	posts, err1 = wmp.ParseFiles("posts")
	dynamicPosts, err2 = wmp.ParseFiles("posts/dynamic")

	if err = cmp.Or(err1, err2); err != nil {
		log.Fatalln(err)
	}

	for k := range tagsMap {
		tagsSorted = append(tagsSorted, k)
	}
	slices.Sort(tagsSorted)

	err1 = sortMarkdownFilesByDate(posts, false)
	err2 = sortMarkdownFilesByDate(dynamicPosts, true)
	if err := cmp.Or(err1, err2); err != nil {
		log.Fatalln(err)
	}

	g := generator()

	if err := g.Generate(); err != nil {
		log.Fatalln(err)
	}
}
