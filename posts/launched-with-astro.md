---
title: "*Launched* with Astro!"
published: 2024-04-04
preamble: Get it? Because Astro's logo is a rocket, and this site is made with Astro! My bad my bad. I know that's not a great way to make a first post with, but I wanted to make that pun so bad. So with the bad joke out of the way, let's get to the heart of the matter.
author: Zill_Laiss
tags:
  - astro
  - go
  - svelte
  - coding
---

## A blog site
And not just an ordinary blog site, it's **MY** blog site! It was build from scratch with my own tears and  blood, worked day and night, with joy and sorry! Awesome right!?

Okay, maybe not that great. Like I said, this site is made with Astro, a great framework created specifically for developing content-driven websites, so in reality it was not hard (plus I took several shortcuts by using available tutorials on the web, but don't tell anybody, okay?). 

At this point, it might seems to you that I love Astro so much since I praise it so many times in that short while. That, kids, is not wrong, and I have so many reasons why.

### Why Astro?

First, let's rewind a little bit. I was planning to make this site with the most loved JS framework according to Stackoverflow: Svelte. And yeah, I actually loved it and got so far in building the site. There's just this one problem though: I overengineered it.

You see, I was only planning to build a simple blog, and by simple, I mean *simple*. 

I don't need anything fancy as a feature for the site. I don't need any complex configuration and rendering techniques just to render my contents. And yet, all I did was the opposite of those. 

Features after features, package after package... At one point, I used Supabase just to store my blogs, and from there, came the authentication along with complete authorization features for the whole site! *Just, to make, a **blog***!

Svelte, or Sveltekit to be precise, is more suitable for building an SPA-like website. Technically, you can build a similar static website with it, but it's hard to resist the temptation of adding unnecessary features. So, off with Svelte.

Then I switched to Go, a.k.a. Golang. Why Go? Well first, it's my go-to programming language for everything. Second, I was developing a desktop app with it concurrently with the blog. I also wanted to try the net/http package + HTMX combination, as it got pretty much hyped in recent times.

Anyway, I started building the website from scratch. Man, it was awesome. I loved every second of it. It's simple, fast, has type safety, and efficient, everything I ever wanted when developing a website. One interesting thing is that, as a result of using HTMX in the project, I now fully comprehend the benefit of both SSR and CSR, when to use and combine them; so I added Alpine too in the end. I want to delve into this topic deeper, but let's not derail and put it aside for another time. Anyway, I loved it and successfully built a skeleton for the site, and you can actually see it [here](http://152.69.212.110/). 

So where's the problem, you ask?

(un)fortunately, I realized something before I continue any further progress: too much hassle for a simple blog.

Again, let's rewind a little bit to before I deployed the skeleton site. 

The thing about Go is that it compiles into a single binary as it is a compiled language. Not to remind you again that I used HTMX, a server-side handling library, so while what I wanted to build was mostly a static site, using a static site only provider would be out of question. I needed a proper server, somewhere.

Not a big problem. Buying a server and host it locally would be a little too far just for a blog, but I can use hosting provider somewhere.

But I came into another obstacle: where to host it?

Now there's a lot of hosting provider, Hetzner, Linode, Google Cloud, you name it. But for some reason, they kept rejecting my registration. I researched more about the matter but it seems like these three have had a problem with scammers using their sites, so they limit their registration scope. How? By blocking a little bit of suspiciousness, like a slightly different spelling in account name vs id card, by not accepting some cards, and one of them just blocked the entire country, *wow*! 

Well, that sucks, I said. Now the probability of me finding a provider was slimmer, as I found out that more sites apply this kind of security. It's not easy for developer who lives outside Europe or the USA, man (At this time, it did not occur to me to host at a local web hosting provider. But I checked out earlier that they mostly only accept a static website hosting, so... yeah).

I considered an alternative to containerize it, but the sites which provide the ability to host a container were either too expensive or had the same problem as before. I did find some accepting ones though, like Digitalocean and Oracle, and I did deploy the skeleton site for testing purpose to Oracle, as you can see at the link above. But at this point, I realized how much of a hassle developing a containerized app is....

After that, I became a little bit frustrated, so I searched for another alternative and the best method to deploy a static site: a static site generator. 

You might ask me a question, why not just using a static site generator from the start if what you really wanted is a static site? That way you could've avoid all of this madness. Well, the thing is, it's correct that right now I only need a static site, but there's a possibility for me to use the site for something else. Yeah...just, a *possibility*. Unfortunately on the way, I fell into a trap hole for all developers out there: premature optimization. Learn from me, kids.

Anyway, I started to check some popular static site generators. I already knew about Hugo for a long time, as it is built with Go. The syntax is familiar for me and it comes with a lot of advanced configurations that you can use to customize your site. Pretty solid choice, but a little bit too complex? It even has Go-styled type safety for every of its features.

Then I checked Astro as it is pretty much hyped in recent times and I was curious of what the fuse is all about, and I don't need to explain how my initial reaction was, right? Everything was so easy! I love that you can just slap JS code you already familiar with. You don't even need to code well as it is a static site generator, as it would not even matter since Astro strips all the JS in the end. Just think about the HTML syntax and CSS. But hey, you need some JS? No worries, you can use its amazing Islands architecture and use the UI framework component you love! Oh you need a little bit of SSR? You can opt-in for a hybrid rendering and just specify a page that needs SSR, and everything else of the site is still static! What's that? You don't want Astro syntax or import a UI framework? You just want to use vanilla JS? Again, Astro covers it for you! Vanilla JS is fully compatible with Astro, unlike some other popular frameworks out there. You can use a client-side JS library like Alpine pretty easily! Man the list just goes on and on. It surprised me that a framework can have this much features at the same time, and still maintains a fast and great DX.

So in the end, the answer is obvious. I needed to ship it fast since all the frustration I'd got before, and I already love Astro ever since I took the first step into it, so there you go. We got to the present day.
## Project tracker
Aside from blogging, this site is used for tracking projects I'm working on. I realize that I could have used Github or similar for it, but there are things that I can't do or just prefer to be in my personal site instead. 

There's no special page for it yet, but I'm planning to create a list of my projects later.
## Playground

Now, this is the fun part. I always love when people test a feature directly ~~on prod~~ on the net before they publish it. Those beta-testing phase is always fun and I want to do the same. It might be not how the final feature would feel or look, but you will find that the foundation is always from there.

I made a semi-hidden page for it, check it out if you want [here](/playground).
## Portfolio?
Maybe, I'm not sure. Other people do it, but I'm not very keen of the idea. I mainly code as a hobby and I don't see a strong reason why I need to write an portfolio. The good ol' PDF is enough, lol.

## Conclusion

I spent too much time writing about Astro than I did with other things in this post, but no ragretz. Either way, It was a great experience creating my first blog site from scratch, and I'm excited to see more of what the future holds.
