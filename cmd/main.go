package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Contact struct {
	Name  string
	Email string
}

type Data struct {
	Contacts []Contact
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			{"John Doe", "john@gmail.com"},
			{"Jane Doe", "jane@gmail.com"},
		},
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	data := newData()
	e.GET("/", func(c echo.Context) error {
		fmt.Print(data.Contacts)
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		data.Contacts = append(data.Contacts, Contact{name, email})
		return c.Render(200, "index", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
