package template

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	fmt.Println("NewTemplate!")
	return &Template{
		// public/views以下の全てのファイルを取り込み、キャッシュ
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
