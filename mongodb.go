package main
import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Person struct {
    Name string
    Age  int
    Phone string
    Detail Details
}

type Details struct {
    Fname string
    Lname string
    Email string
}

func main() {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonicbehavior.
    // session.SetMode(mgo.Monotonic, true)

    c := session.DB("test").C("people")
    err = c.Insert(&Person{"Ale", 11, "+55 53 8116", Details{"ahmed", "adams", "a@a.com"}},
                    &Person{"jameel", 12, "+9055 666 777", Details{"ahmed", "adams", "a@a.com"}},
                    &Person{"adil", 13, "+9055 666 888", Details{"ahmed", "adams", "a@a.com"}},
                    &Person{"foad1", 14, "+9055 666 999", Details{"ahmed", "adams", "a@a.com"}})

    if err != nil {
        panic(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "foad1"}).One(&result)
    if err != nil {
        log.Fatal(err)
    }


    count, err := c.Count()
    if err != nil {
        fmt.Println("err is ", err)
    }
    fmt.Println("count of collections is : ", count)

    fmt.Println("name:", result.Name)
    fmt.Println("Phone:", result.Phone)
    fmt.Println("age:", result.Age)
    fmt.Println("fname: ", result.Detail.Fname)
    fmt.Println("email: ", result.Detail.Email)

}
