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
