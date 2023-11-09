package main

import (
	"github.com/denisemignoli/restful-api-with-go-and-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run("localhost:8080")
}
