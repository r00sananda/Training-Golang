package main

import (
	"log"
	"training-golang/DAY-5/entity"
	"training-golang/DAY-5/handler"
	"training-golang/DAY-5/repository/slice"
	"training-golang/DAY-5/router"
	"training-golang/DAY-5/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
