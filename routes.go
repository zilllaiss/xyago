package main

import (
	"context"
	"github.com/zilllaiss/fest"
	"github.com/zilllaiss/fest/temfest"
	"xyago/types"
	"xyago/views"

	"github.com/a-h/templ"
)

func generator() *fest.Generator {
	g := fest.NewGenerator(context.Background(), "Zill_Laiss' blog", nil)

	g.CopyDir("assets", "")
	g.CopyFile("favicon.png", "")
	g.CopyFile("misc/README.md", "")

	g.Head.Add(
		temfest.ImportStyle("/assets/styles/global.css"),
		temfest.ImportIcon("/favicon.png", "image/png"),
	)

	g.AddRouteFunc("/", index)
	g.AddRoute("/404.html", views.NotFound())
	g.AddRoute("/about", views.About()).SetTitle("About")
	g.AddRoute("/tags", views.Tags(tagsSorted))

	fest.NewRoutes("/tags/{s}", tagsSorted).
		SetTitle("{s}").AddToGenerator(g, tagsFn)

	fest.NewRoutesT("/posts/{s}", posts).
		SetTitle("{s}").AddToGenerator(g, postsFn)

	fest.NewPaginatedRoutes("/blog/{s}", posts, 5).
		SetTitle("Blogs - page {s}").AddToGenerator(g, blogs)

	return g
}

func index(ctx context.Context) (templ.Component, error) { return views.Index(posts), nil }

func blogs(ctx context.Context,
	rp *fest.RoutesParam[*fest.Pagination[types.Markdown]],
) (templ.Component, error) {
	return views.Blog(tagsMap, tagsSorted, *rp.GetItem()), nil
}

func tagsFn(ctx context.Context, rp *fest.RoutesParam[string]) (templ.Component, error) {
	return views.Tag(tagsMap), nil
}

func postsFn(ctx context.Context, rp *fest.RoutesParam[types.Markdown]) (templ.Component, error) {
	post := rp.GetItem()
	rp.SetSlug(post.Slug)
	return views.Post(post), nil
}
