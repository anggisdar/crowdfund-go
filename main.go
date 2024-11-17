package main

import (
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

	userRepository := user.NewRepository(dbgo)
	user := user.User{
		Name: "Test simpan",
	}

	userRepository.Save(user)
}
