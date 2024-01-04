package routes

import (
	"github.com/denisemignoli/restful-api-with-go-and-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	albumController := controllers.NewAlbumController()

	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbumByID)
	router.POST("/albums", albumController.PostAlbums)
}
