package post

import (
	"bytes"
	"html/template"
	"io/fs"
	"maps"
	"slices"
	"time"

	"github.com/obzva/dngyng1000/internal/ui"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

type Meta struct {
	Slug        string    `yaml:"slug"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	CreatedAt   time.Time `yaml:"created_at"`
	UpdatedAt   time.Time `yaml:"updated_at"`
}

type Post struct {
	Meta
	Body template.HTML
}

type Map map[string]*Post

func NewMap() (Map, error) {
	pm := make(Map)

	posts, err := fs.Glob(ui.Files, "posts/*.md")
	if err != nil {
		return nil, err
	}

	// initialize the markdown parser
	mdParser := goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{}))

	for _, post := range posts {
		// read source file
		src, err := fs.ReadFile(ui.Files, post)
		if err != nil {
			return nil, err
		}

		// parse markdown
		var out bytes.Buffer
		ctx := parser.NewContext()
		if err := mdParser.Convert(src, &out, parser.WithContext(ctx)); err != nil {
			return nil, err
		}

		// extract m
		var m Meta
		if err := frontmatter.Get(ctx).Decode(&m); err != nil {
			return nil, err
		}

		// register Post struct to the map
		p := Post{
			Meta: m,
			Body: template.HTML(out.String()),
		}
		pm[p.Slug] = &p
	}

	return pm, nil
}

func (m Map) Slices() []*Post {
	it := maps.Values(m)

	sorter := func(a, b *Post) int {
		return a.CreatedAt.Compare(b.CreatedAt)
	}

	return slices.SortedFunc(it, sorter)
}
