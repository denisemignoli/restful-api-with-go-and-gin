package controllers

import (
	"github.com/denisemignoli/restful-api-with-go-and-gin/models"
	"github.com/denisemignoli/restful-api-with-go-and-gin/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumRepository *repositories.AlbumMySQLRepository
}

func NewAlbumController() *AlbumController {
	return &AlbumController{
		AlbumRepository: repositories.NewAlbumMySQLRepository(),
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
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id format"})
		return
	}

	// Passe o id como int64 para a função GetAlbumByID
	album, err := ac.AlbumRepository.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
