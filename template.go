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

    // access to field in struct by point like this : {{.Feild}} 
    t := template.New("text 0")
    t, _ = t.Parse("hello {{.Name}}! Age is {{.Age}}.\n")
    err := t.Execute(os.Stdout, p)

    if err != nil {
        fmt.Println("ther was an error:", err)
    }

    name := "adam"
    // access to value of varialbl just by point like this : {{.}} 
    t1 := template.New("example 1")
    t1, _ = t1.Parse("Hi! {{if `true`}} {{.}} {{else}} Evryone {{end}}")
    err = t1.Execute(os.Stdout, name)

    if err != nil {
        fmt.Println("ther was an error:", err)
    }
}
