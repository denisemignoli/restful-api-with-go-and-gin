package controllers

import (
	"net/http"

	"github.com/denisemignoli/restful-api-with-go-and-gin/models"
	"github.com/denisemignoli/restful-api-with-go-and-gin/repositories"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumRepository *repositories.AlbumRepository
}

func NewAlbumController() *AlbumController {
	return &AlbumController{
		AlbumRepository: &repositories.AlbumRepository{},
	}
}

func (ac *AlbumController) GetAlbums(c *gin.Context) {
	albums := ac.AlbumRepository.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func (ac *AlbumController) PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
	}
	ac.AlbumRepository.SaveAlbums(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (ac *AlbumController) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := ac.AlbumRepository.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
