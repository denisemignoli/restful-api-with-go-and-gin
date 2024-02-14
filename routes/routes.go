package routes

import (
	"github.com/denisemignoli/restful-api-with-go-and-gin/controllers"
	"github.com/denisemignoli/restful-api-with-go-and-gin/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// para mysql
	albumController := controllers.NewAlbumController(repositories.NewAlbumMySQLRepository())

	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbumByID)
	router.POST("/albums", albumController.PostAlbums)
}
