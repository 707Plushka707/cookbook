package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	// "gopkg.in/mgo.v2"
	"github.com/labstack/echo/v4"
)

// template ============================
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("assets/form.html")),
}

// ===================================
type User struct {
	Title  string `json:"title" xml:"email" form:"email" query:"email"`
	Articl string `json:"articl" xml:"name" form:"name" query:"name"`
}

func getData(c echo.Context) error {

	return c.Render(http.StatusOK, "form.html", nil)
}

func postData(c echo.Context) error {
	var u User
	u.Title = c.FormValue("title")
	u.Articl = c.FormValue("articl")

	fmt.Printf("title is : %s,  articl is : %s", u.Title, u.Articl)

	return c.JSON(200, u)
}

func main() {
	e := echo.New()
	e.Renderer = t
	e.GET("/", getData)
	e.POST("/", postData)
	fmt.Println(e.Start(":8080"))
}
