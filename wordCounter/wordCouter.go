package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var words []string
	var line []string

	for scanner.Scan() {
		line = strings.Split(scanner.Text(), " ")
		for _, w := range line {
			//words = strings.Trim(w," .,!?-\"")
			words = append(words, strings.Trim(w, " .,!?-\""))
		}

	}

	nwords := make(map[string]int)
	// count words
	for _, w := range words {
		nwords[w] = nwords[w] + 1
	}
	for w, n := range nwords {
		if n > 3 {
			fmt.Println(w, "   ", n)
		}
	}
	fmt.Println("Done")
}
