package main
import (
    "strconv"
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main() {
    rows := [][]string{
        {"Name", "City", "Language","year"},
        {"Pinky", "london", "Python", "12"},
        {"Nicky", "Paris", "Golang", "13"},
        {"Micky", "Tokyo", "Php", "15"},
    }
    csvfile, err := os.Create("test.csv")
    checkError("failed creating file: ", err)


    csvwriter := csv.NewWriter(csvfile)

    for _, row := range rows {
        _ = csvwriter.Write(row)
    }

    csvwriter.Flush()
    csvfile.Close()
    year := len(rows[0])-1
    years := 0
    for _, r := range rows[1:] {
        y, _ := strconv.Atoi(r[year])
        years += y
    }
    fmt.Println(years)
}

func checkError(str string, err error) {
    if err != nil {
        log.Fatalln(str, err)
        fmt.Println()
        return
    }
}
