package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var a = "hi"
	switch a {
	case "hi":
		fmt.Println("a is hi")
	case "jawad":
		fmt.Println("a is jawad")
	default:
		fmt.Println("a is something esle")
	}
}
