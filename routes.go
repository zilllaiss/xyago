package main

import (
	"context"

	"xyago/types"
	"xyago/views"

	"github.com/zilllaiss/fest"
	"github.com/zilllaiss/fest/markdown"
	"github.com/zilllaiss/fest/temfest"

	"github.com/a-h/templ"
)

const postStyling = "/assets/styles/post.css"

func generator() *fest.Generator {
	g := fest.NewGenerator(context.Background(), "Zill_Laiss' blog", &fest.GeneratorConfig{
		BaseConfig: temfest.BaseConfig{Lang: "en"},
	})

	g.CopyDir("assets", "")
	g.CopyFile("favicon.png", "")
	g.CopyFile("misc/README.md", "")

	g.HeadBody.Head(
		temfest.ImportStyle("/assets/styles/global.css"),
		temfest.ImportIcon("/favicon.png", "image/png"),
	)

	g.AddRouteFunc("/", index).HeadBody.Head(temfest.ImportStyle(postStyling))
	g.AddRoute("/404.html", views.NotFound())
	g.AddRoute("/about", views.About()).SetTitle("About")
	g.AddRoute("/tags", views.Tags(tagsSorted)).SetTitle("Tags")

	tagRoutes := fest.NewRoutes("/tags/{s}", tagsSorted)
	tagRoutes.HeadBody.Head(temfest.ImportStyle(postStyling))
	tagRoutes.AddToGenerator(g, tagsFn)

	postRoutes := fest.NewRoutesT("/posts/{s}", posts)
	postRoutes.HeadBody.Head(temfest.ImportStyle(postStyling))
	postRoutes.AddToGenerator(g, postsFn)

	blogPages := fest.NewPaginatedRoutes("/blog/{s}", posts, 5).
		SetTitle("Blogs - page {s}")
	blogPages.HeadBody.Head(temfest.ImportStyle(postStyling))
	blogPages.AddToGenerator(g, blogs)

	return g
}

func index(ctx context.Context) (templ.Component, error) { return views.Index(posts), nil }

func blogs(ctx context.Context,
	rp *fest.RouteParam[*fest.Pagination[*markdown.MarkdownData]],
) (templ.Component, error) {
	return views.Blog(tagsMap, tagsSorted, *rp.GetItem()), nil
}

func tagsFn(ctx context.Context, rp *fest.RouteParam[string]) (templ.Component, error) {
	return views.Tag(tagsMap), nil
}

func postsFn(ctx context.Context, rp *fest.RouteParam[*markdown.MarkdownData]) (templ.Component, error) {
	post := rp.GetItem()
	fm := &types.Frontmatter{}

	if err := post.GetFrontmatter(fm); err != nil {
		return nil, err
	}

	rp.SetSlug(post.Slug)
	rp.SetTitle(fm.Title)

	return views.Post(post), nil
}
