package views

import (
	"embed"
	"html/template"
	"io"
)

//go:embed templates/*.html
var TS embed.FS

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any) error {
	templ, err := template.ParseFS(TS, "templates/"+name, "templates/base.html")
	if err != nil {
	  return err
	}
  
	return templ.ExecuteTemplate(w, "base", data)
  }