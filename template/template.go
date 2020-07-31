package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	fmt.Printf("%%\n")
	p := Person{Name: "admin", Age: "34"}

	// access to field in struct by point like this : {{.Feild}}
	t := template.New("text 0")
	t, _ = t.Parse("hello {{.Name}}! Age is {{.Age}}.\n")
	err := t.Execute(os.Stdout, p)
	checkeErr("err when parsing struct", err)

	name := "adam"
	// access to value of varialbl just by point like this : {{.}}
	t1 := template.New("example 1")
	t1, _ = t1.Parse("Hi! {{if `true`}} {{.}} {{else}} Evryone {{end}}")
	err = t1.Execute(os.Stdout, name)
	checkeErr("err when parsing variable", err)

	slc := []int{2, 3, 4, 5, 6}
	t2 := template.New("example 1")
	t2, _ = t1.Parse("range  {{range .}} {{.}} {{end}}")
	err = t2.Execute(os.Stdout, slc)
	checkeErr("err when parsing slice", err)
}

func checkeErr(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(0)
	}
}
