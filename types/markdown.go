// Package types provides types used globally in the project.
// This package should not import anything from the project itself.
package types

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
	Filename string `yaml:"filename"` // relative from /assets/images
	Alt      string `yaml:"alt"`
	Hover    string `yaml:"hover"` // title attribute
}
