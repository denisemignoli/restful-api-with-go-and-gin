package controllers

import (
	"net/http"

	"github.com/denisemignoli/restful-api-with-go-and-gin/models"
	"github.com/denisemignoli/restful-api-with-go-and-gin/repositories"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albumRepository := &repositories.AlbumRepository{}
	albums := albumRepository.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	albumRepository := &repositories.AlbumRepository{}

	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
	}
	albumRepository.PostAlbums(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albumRepository := &repositories.AlbumRepository{}

	album, err := albumRepository.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
