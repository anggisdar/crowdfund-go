package main

import (
	"fmt"
	"log"
	"webfunding/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:05072001@tcp(localhost:3306)/dbgo"
	dbgo, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database conected")

	var users []user.User

	dbgo.Find(&users)

	length := len(users)
	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}
}
