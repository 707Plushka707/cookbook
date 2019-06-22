package main
import "fmt"

type Person struct{
    Name string
    age int
}

func main() {
    person := &Person{"ahmed", 16}
    fmt.Println(person.Name)
}
