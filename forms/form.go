package main

import (
<<<<<<< HEAD
	"fmt"
	"html/template"
	"net/http"
	"os"
=======
    "html/template"
    "net/http"
    "fmt"
    "os"
>>>>>>> 3bd77332b2e755fc6ed6e36c775f0d23fbb7ab00

	"gopkg.in/mgo.v2"
)

type Book struct {
<<<<<<< HEAD
	Name    string
	Subject string
	Author  string
=======
    Name   string
    Subject string
    Author string
>>>>>>> 3bd77332b2e755fc6ed6e36c775f0d23fbb7ab00
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	CheckeErr("err in create session: maybe mongodb not active", err)

	defer session.Close()

	tmpl := template.Must(template.ParseFiles("asset/form.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
	    tmpl.Execute(w, nil)
	    return
	}

		c := session.DB("test").C("Book")

<<<<<<< HEAD
		book := Book{
			Name:    r.FormValue("name"),
			Subject: r.FormValue("subject"),
			Author:  r.FormValue("author"),
		}
		err = c.Insert(&Book{book.Name, book.Subject, book.Author})

		// do something with details
		fmt.Println(book.Name)
		fmt.Println(book.Subject)

		tmpl.Execute(w, struct{ Ok bool }{true})
	})

=======
	book := Book{
		Name:   r.FormValue("name"),
		Subject: r.FormValue("subject"),
	        Author: r.FormValue("author"),
	}
        err = c.Insert(&Book{book.Name, book.Subject, book.Author})

	// do something with details
        fmt.Println(book.Name)
        fmt.Println(book.Subject)

        tmpl.Execute(w, struct{Ok bool}{true})
	
>>>>>>> 3bd77332b2e755fc6ed6e36c775f0d23fbb7ab00
	http.ListenAndServe(":8080", nil)
}

func CheckeErr(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(0)
	}
}
