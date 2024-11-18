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
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Tes simpan dari service"
	userInput.Email = "anggsidar@gmail.com"
	userInput.Ocuppation = "anak band"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}

//inputnya dari user
//handler : mapping input dari user > struct input
//service : melakukan mapping dari struct entity.go ke struct user
//repository
//db
