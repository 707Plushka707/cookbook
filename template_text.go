package main
import (
    "text/template"
    "os"
)
func main() {
    t := template.New("template test")
    t = template.Must(t.Parse("This is just static text. \n {{\"this is nothing\"}} {{`\nSo this version quote`}}\n\n\n"))
    t.Execute(os.Stdout, nil)

    tEmpty := template.New("template test")
    tEmpty = template.Must(tEmpty.Parse("no Empty pipeline if demo: \n {{if `anything`}} Will  print. {{end}}")) // empty empty folowing if
    tEmpty.Execute(os.Stdout, nil)

    tWithVal := template.New("template test")
    tWithVal = template.Must(tWithVal.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}")) // empty empty folowing if
    tWithVal.Execute(os.Stdout, nil)

    tIfElse := template.New("template test")
    tIfElse = template.Must(tIfElse.Parse("if else demo: {{if ``}} print if part.{{else}} \nprint else part {{end}}")) // empty empty folowing if
    tIfElse.Execute(os.Stdout, nil)

    t1, _ := template.New("test").Parse("{{with `hello`}}{{.}}{{end}}\n")
    t1.Execute(os.Stdout, nil)

    // when nested, the dot takes the value according to closet scope.
    t2, _ := template.New("test").Parse("{{with `hello`}}{{.}}{{with `Admin`}}{{.}}{{end}}{{end}}!\n")
    t2.Execute(os.Stdout, nil)

    var1 := template.Must(template.New("var1").Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
    var1.Execute(os.Stdout, nil)

    var2 := template.Must(template.New("var2").Parse("{{with $x3 := `welcome`}}{{$x3}}{{end}}!\n"))
    var2.Execute(os.Stdout, nil)

    var3 := template.Must(template.New("var3").Parse("{{with $_3 := `salam`}}{{$_3}}{{end}}!\n"))
    var3.Execute(os.Stdout, nil)

    tfunc := template.New("tempFunc")
    tfunc = template.Must(tfunc.Parse("{{with $x := `hello`}}{{printf `%s %s` $x  `Admin`}}{{end}}!\n"))
    tfunc.Execute(os.Stdout, nil)

}
