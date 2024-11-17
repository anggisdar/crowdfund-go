package main

import (
	"log"
	"net/http"
	"webfunding/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 	dsn := "root:05072001@tcp(localhost:3306)/dbgo"
	// 	dbgo, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}

	// 	fmt.Println("Connection to database conected")

	// 	var users []user.User
	// 	dbgo.Find(&users)

	// 	for _, user := range users {
	// 		fmt.Println(user.Name)
	// 		fmt.Println(user.Email)
	// 	}
	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:05072001@tcp(localhost:3306)/dbgo"
	dbgo, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	dbgo.Find(&users)

	c.JSON(http.StatusOK, users)
}
