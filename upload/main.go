package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	// read form fields
	name := c.FormValue("name")

	// multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("public/" + file.Filename)
		if err != nil {
			return err
		}

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Upload successfully %d files with fields name=%s </p><img src='%s'/>", len(files), name, files[0].Filename))
}
func main() {
	e := echo.New()

	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)
	fmt.Println(e.Start(":1234"))

}
