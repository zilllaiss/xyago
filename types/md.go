// Package types provides types used globally in the project.
// This package should not import anything from the project itself.
package types

type Markdown struct {
	Slug        string
	Frontmatter Frontmatter
	Content     string
	TOC         string
}

type Frontmatter struct {
	Title       string    `yaml:"title"`
	Author      string    `yaml:"author"`
	Public      bool      `yaml:"public"`
	PublishedAt string    `yaml:"published"`
	UpdatedAt   string    `yaml:"updated"`
	Preamble    string    `yaml:"preamble"`
	Tags        []string  `yaml:"tags"`
	Eyecatch    *Eyecatch `yaml:"eyecatch"`
}

type Eyecatch struct {
	Path string `yaml:"path"`
	Alt  string `yaml:"alt"`
}
