package main

import (
	"log"
	"training-golang/DAY-2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(r)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}
