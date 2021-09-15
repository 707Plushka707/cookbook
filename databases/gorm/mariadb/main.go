package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	//gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phon       string `json:"phon"`
	Avatarlink string `json:"avatarlink"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/sooq?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Users{})

	// Create
	db.Create(&Users{Username: "jawad abkar", Password: "123456", Email: "ab@gorm.com", Phon: "05344667788", Avatarlink: "www.image.com/myimge2.png"})

	// Read
	var client Users
	//db.First(&client, 1)                     // find product with id 1
	db.First(&client, "userid = ?", 1) // find product with code l1212
	fmt.Println(client)

	// Update - update product's price to 2000
	//db.Model(&client).Update("username", "admin")

	// Delete - delete product
	//db.Delete(&client)
}
