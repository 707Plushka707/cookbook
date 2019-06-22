// this is an example of how to connect mariadb with golang apps
// mariadb is like mysql . and this example is usefall with both database

package main
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("connecting with database")
    // my database name is db1; you cna use any name
    // use your name password and youl local database

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
    if err != nil {
        fmt.Println(err)
    }


    defer db.Close()
    fmt.Println("database connected")

    insert , err := db.Query("INSERT INTO test VALUES (1, 'first_test_db')")
    if err != nil {
        fmt.Printf("error in Query insert %v\n", err.Error())
    }

    //don't ignore closing Queries if you are using transactions
    defer insert.Close()

}
