package main
import (
    "fmt"
    "net/http"
)

// both type emplements serveHttp

type (
    String string
    Struct struct {
        Greating string
        Punct string
        Who string
    }
)

// testing

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", s)
}

// another testing
func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s%s %s", s.Greating, s.Punct, s.Who)
}

func main() {
    // http.Handle calls
    http.Handle("/string", String("I'am a frid kno"))
    http.Handle("/truct", &Struct{
        Greating: "Hello struct",
        Punct: ",",
        Who: "World!",
    })
    http.Handle("/", &Struct{"Hello /", ",", "World"})
    fmt.Println("Listining on port :8080")
    http.ListenAndServe("localhost:8080", nil)
}
