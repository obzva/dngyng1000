package template

import (
	"bytes"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/obzva/dngyng1000/internal/ui"
)

type Cache map[string]*template.Template

func NewCache() (Cache, error) {
	tc := make(Cache)

	pages, err := fs.Glob(ui.Files, "tmpls/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFS(ui.Files, "tmpls/base.tmpl", page)
		if err != nil {
			return nil, err
		}

		tc[name] = ts
	}

	return tc, nil
}

func (c Cache) Render(logger *slog.Logger, w http.ResponseWriter, r *http.Request, tmplFileName string, data any) {
	ts, ok := c[tmplFileName]
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// buffer for trial render
	// it helps us to catch runtime error
	// if trial render onto this buffer succeed then we render the content onto http.ResponseWriter
	var b bytes.Buffer

	if err := ts.ExecuteTemplate(&b, "base", data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if _, err := b.WriteTo(w); err != nil {
		logger.Error(err.Error())
	}
}
