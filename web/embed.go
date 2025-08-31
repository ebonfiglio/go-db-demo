package web

import (
	"embed"
	"html/template"
	"io/fs"
)

//go:embed templates templates/organizations
var rawFS embed.FS

var FS, _ = fs.Sub(rawFS, "..")

func Parse() *template.Template {
	t := template.New("")

	patterns := []string{
		"templates/*.html",
		"templates/organizations/*.html",
	}

	for _, pat := range patterns {
		matches, _ := fs.Glob(FS, pat)
		if len(matches) == 0 {
			continue
		}
		template.Must(t.ParseFS(FS, pat))
	}

	// Optional fallback to avoid panic if no files exist yet
	if t.Tree == nil || len(t.Templates()) == 0 {
		template.Must(t.New("fallback").Parse(`ok`))
	}

	return t
}
