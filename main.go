package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.jBvaWZVpqBF9oL0DTMp31fQtpJuk1D9sihuBvxz7U0c")
	// if err != nil {
	// 	fmt.Println("Error")
	// 	fmt.Println("Error")
	// 	fmt.Println("Error")
	// }

	// if token.Valid {
	// 	fmt.Println("Valid")
	// 	fmt.Println("Valid")
	// 	fmt.Println("Valid")
	// } else {
	// 	fmt.Println("Invalid")
	// 	fmt.Println("Invalid")
	// 	fmt.Println("Invalid")
	// }

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run(":8070")
}
