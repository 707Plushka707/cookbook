package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	// init some data
	welcome := Welcome{"Anonymoce", time.Now().Format(time.Stamp)}

	// handle templates files
	templates := template.Must(template.ParseFiles("templ/welcome.html"))

	// handle css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle request
	//http.Handle("/static/", http.FileServer(http.Dir("static/stylesheets/")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server start at :9000", http.ListenAndServe(":9000", nil))

}
