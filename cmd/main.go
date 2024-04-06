package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

func (t *Templates) Render(writter io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(writter, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	count := Count{Count: 0}

	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/css", "css")

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "index", count)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
