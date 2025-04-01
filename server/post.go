package server

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	goldmarkMeta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

var postMap = NewPostList()

type Post struct {
	ID, Title, Description string
	Date                   time.Time
	Body                   template.HTML
}

type PostList struct {
	// map of post id to post
	m map[string]*Post
	// slice of posts
	s []*Post
}

func (pm *PostList) Get(id string) (*Post, error) {
	post, ok := pm.m[id]
	if !ok {
		return nil, fmt.Errorf("post with id %s not found", id)
	}
	return post, nil
}

func NewPostList() *PostList {
	dirName := "posts"
	dir, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	pl := &PostList{
		m: make(map[string]*Post),
		s: make([]*Post, 0),
	}

	for _, entry := range dir {
		p, err := getPost(path.Join(dirName, entry.Name()))
		if err != nil {
			panic(err)
		}

		pl.s = append(pl.s, p)
		pl.m[p.ID] = p
	}

	return pl
}

// open postfile from fileSystem and return new Post struct
func getPost(fileName string) (*Post, error) {
	postFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

// make new Post struct from io.Reader
func newPost(postFile io.Reader) (*Post, error) {
	// read all bytes from postFile
	b, err := io.ReadAll(postFile)
	if err != nil {
		return nil, err
	}

	// parse bodyBuffer and meta from md file
	var bodyBuffer bytes.Buffer
	pc := parser.NewContext()
	md := goldmark.New(goldmark.WithExtensions(goldmarkMeta.Meta))
	if err := md.Convert(b, &bodyBuffer, parser.WithContext(pc)); err != nil {
		return nil, err
	}

	meta := goldmarkMeta.Get(pc)

	title, err := assertString(meta["title"])
	if err != nil {
		return nil, err
	}

	id := strings.ToLower(title)
	id = strings.ReplaceAll(id, " ", "-")
	id = url.QueryEscape(id)

	description, err := assertString(meta["description"])
	if err != nil {
		return nil, err
	}

	dateString, err := assertString(meta["date"])
	if err != nil {
		return nil, err
	}

	date, err := parseDate(dateString)
	if err != nil {
		return nil, err
	}

	return &Post{
		ID:          id,
		Title:       title,
		Description: description,
		Date:        date,
		Body:        template.HTML(bodyBuffer.String()),
	}, nil
}

func assertInt(v interface{}) (int, error) {
	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("expected int, got %T", v)
	}
	return i, nil
}

func assertString(v interface{}) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %T", v)
	}
	return s, nil
}

func parseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}
