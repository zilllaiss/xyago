package main

import (
	"context"

	"xyago/types"
	"xyago/views"

	"codeberg.org/zill_laiss/fest"
	"codeberg.org/zill_laiss/fest/markdown"
	"codeberg.org/zill_laiss/fest/temfest"

	"github.com/a-h/templ"
)

const postStyling = "/assets/styles/post.css"

func generator() *fest.Generator {
	siteName := "Zill_Laiss' blog"
	suffix := " - " + siteName

	g := fest.NewGenerator(context.Background(), siteName, nil)
	g.SetLanguage("en")

	g.CopyDir("assets", "")
	g.CopyFile("favicon.png", "")
	g.CopyFile("misc/README.md", "")

	g.AppendToHead(
		temfest.ImportStyle("/assets/styles/global.css"),
		temfest.ImportIcon("/favicon.png", "image/png"),
	)

	g.AddRouteFunc("/", index).
		SetTitle(siteName).
		AppendToHead(temfest.ImportStyle(postStyling))

	g.AddRoute("/404.html", views.NotFound()).SetTitle("Not found")
	g.AddRoute("/about", views.About()).SetTitle("About" + suffix)
	g.AddRoute("/tags", views.Tags(tagsSorted)).SetTitle("Tags" + suffix)

	fest.NewRoutes("/tags/{s}", tagsSorted).
		SetTitle("{s}"+suffix).
		AppendToHead(temfest.ImportStyle(postStyling)).
		AddToGenerator(g, tagsFn)

	fest.NewRoutesT("/posts/{s}", posts).
		AppendToHead(temfest.ImportStyle(postStyling)).
		AddToGenerator(g, postsFn(suffix))

	fest.NewPaginatedRoutes("/blog/{s}", posts, 5).
		SetTitle("Blogs - page {s}"+suffix).
		AppendToHead(temfest.ImportStyle(postStyling)).
		AddToGenerator(g, blogs)

	return g
}

func index(ctx context.Context) (templ.Component, error) { return views.Index(posts), nil }

func blogs(ctx context.Context,
	rp *fest.RouteParam[*fest.Pagination[*markdown.MarkdownData]],
) (templ.Component, error) {
	return views.Blog(tagsMap, tagsSorted, *rp.GetItem()), nil
}

func tagsFn(ctx context.Context, rp *fest.RouteParam[string]) (templ.Component, error) {
	tag := rp.GetItem()
	return views.Tag(tag, tagsMap), nil
}

func postsFn(suffix string) func(ctx context.Context, rp *fest.RouteParam[*markdown.MarkdownData]) (templ.Component, error) {
	return func(ctx context.Context, rp *fest.RouteParam[*markdown.MarkdownData]) (templ.Component, error) {
		post := rp.GetItem()
		fm := &types.Frontmatter{}

		if err := post.GetFrontmatter(fm); err != nil {
			return nil, err
		}

		rp.SetSlug(post.Slug)
		rp.SetTitle(fm.Title + suffix)

		return views.Post(post), nil
	}
}
