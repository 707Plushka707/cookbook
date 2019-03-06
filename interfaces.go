package main

import "fmt"

type great interface {
    hello() string
}

func (g great) hello() {
    fmt.Println("say hello")
}

func main() {
    g := great()
    g.hello()
    fmt.Println("done")
}
