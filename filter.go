package main
import (
    "fmt"
    "bufio"
    "log"
    "os"
)

func main() {
    file, err := os.Open("bash_history")
    if err != nil {log.Fatalf("faild opening file: %s", err)}
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var textlines []string

    for scanner.Scan() {
        textlines = append(textlines, scanner.Text())
    }

    for _, eachline := range textlines {
        fmt.Println(eachline)
    }
}
