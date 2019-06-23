package main
import (
    "fmt"
    "text/template"
    "os"
)

type Person struct {
    Name string
    Age string
}

func main() {
    p := Person{Name: "admin", Age:"34"}
    t := template.New("nonexported template demo")
    t, _ = t.Parse("hello {{.Name}}! Age is {{.Age}}.")
    err := t.Execute(os.Stdout, p)
    if err != nil {
        fmt.Println("ther was an error:", err)
    }
}
