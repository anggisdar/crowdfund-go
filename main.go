package main

import (
	"log"
	"webfunding/auth"
	"webfunding/handler"
	"webfunding/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:05072001@tcp(localhost:3306)/dbgo?parseTime=true"
	dbgo, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var existingUser user.User
	err = dbgo.Where("email = ?", "anggisdar@gmail.com").First(&existingUser).Error
	if err != nil {
		// Password yang di-hash hanya dilakukan saat membuat user baru
		password := "password" // password yang belum di-hash
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err.Error())
		}

		// User baru dengan password yang sudah di-hash
		newUser := user.User{
			Name:         "Anggis",
			Occupation:   "Developer Enthusiast",
			Email:        "anggisdar@gmail.com",
			PasswordHash: string(passwordHash),
			Role:         "user",
		}

		// Simpan user baru
		err = dbgo.Create(&newUser).Error
		if err != nil {
			log.Fatal("Error inserting user: ", err)
		}
		log.Println("User 'anggis' created successfully with hashed password.")
	}

	userRepository := user.NewRepository(dbgo)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	{
		authService.GenerateToken(1001)
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
