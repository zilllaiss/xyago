---
title: Rewritten in Go
published: 2025-07-23
updated: 2025-08-13
preamble: | 
    More than a year ago, I was relly excited to build this site with Astro, a popular static site generator.
    However, that excitement ended up to not last very long.
author: Zill_Laiss
tags:
  - astro
  - coding
  - fest
  - go
  - templ
---

This site is now written with [FEST](https://github.com/zilllaiss/FEST), a minimalistic static-site generator I made.
I want to say that it is _ported_ since the new source code for is one-to-one with original code written in Astro-js in most area, but there are some big changes. 

Most of the change is in the routing. Astro follow a folder-based routing. That means you make a folder with route name yous wish, and in that folder you made a `index.html`. Most of JS frameworks have similar features for routing.

Now, the routing in Go is more similar to router library like Chi (actually one of my inspiration) or Express.js.

Another big change that I had to make is that I need to implemented convert Markdown files myself, although it is easily handled as Go libraries for that already exists.

## The Reasons Why

Every time I revisit my blog, something breaks. Whenever I opened the site, thinking about writing or tweaking something to my site—and instead, I was greeted with warnings, API and documentation changes, and worse—security issues. 
To me it's funny that a static-site generator has security issues, but knowing Astro has a lot more features, I suppose it's natural for it to have them. 

I really wanted to ignore those as I only use it for static-site, but literally npm told me that there are _20 vulnerabilities_.

To be fair, this isn’t strictly Astro’s fault. This is a broader issue with the JavaScript ecosystem—where even doing *nothing* can result in things silently breaking over time. 

It’s hard to stay motivated when you feel like you're constantly fixing instead of building.

I want is stability. I only need a static website. No need for constant upgrades just to keep things from falling apart. That’s why I wrote my own static site generator in Go—using [Templ](https://templ.guide/). It's minimal, fast, and worry-free. I’ve been using Go and Templ intensively for the past two years, and I love the DX that comes along with it.

## Afterthought

Astro is still good, just not for me, It’s well-designed, and for a lot of people it’s the right choice. But it’s made to be a feature-rich platform—a part of the fast-evolving JavaScript world. That means it comes with a long tail of dependencies and expectations for maintenance.

That’s just not what I want right now.

## What’s Next

1. I’m migrating from Sass to plain CSS** 

Sass doesn’t break as often, but it still does occasionally. I also considered Tailwind, but using it with Templ is not a good of an idea. Go build is fast, but if you need to change something in HTML will need you to rebuild the binary. So it can get tedious pretty fast, especially since I'm using older and weaker laptop.

I do use 'hot-reload', which is a combination of Task (aka. go-task) watch feature and live-server, so it's not that bad. However, different with HTML files itself, I often thinker the styling, so having it rebuild the binary every time for CSS change is the last thing I want.

And frankly, I just love a more minimalistic choice when it is make sense to do so (and most of the time do).

2. I’ll finish the site. 

Most of the pages are unfinished. Actually, most of them haven't. The ones that I considered "mostly finished" are the Blog and Post pages.

