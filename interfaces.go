package main

import "fmt"

type great interface {
    hello()
}

func (g great) hello() {
    fmt.Println("say hello")
}

func main() {
    g := great()
    g.hello()
    fmt.Println("done")
    // this commonte for test git commit
    // this comment for test githubrepo
}
