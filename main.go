package main

import (
	"fmt"
	"log"
	"webfunding/handler"
	"webfunding/user"

	"github.com/gin-gonic/gin"
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

	userByEmail, err := userRepository.FindByEmail("anggisdar@gmail.com")
	if err != nil {
		fmt.Println(err.Error())

	}

	if userByEmail.ID == 0 {
		fmt.Println("User tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}

	fmt.Println(userByEmail)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}

//inputnya dari user
//handler : mapping input dari user > struct input
//service : melakukan mapping dari struct entity.go ke struct user
//repository
//db
