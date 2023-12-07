package repositories

import (
	"errors"

	"github.com/denisemignoli/restful-api-with-go-and-gin/models"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type AlbumRepository struct{}

func (a *AlbumRepository) GetAlbums() []models.Album {
	return albums
}

func (a *AlbumRepository) SaveAlbums(newAlbum models.Album) {
	albums = append(albums, newAlbum)
}

func (a *AlbumRepository) GetAlbumByID(id string) (models.Album, error) {
	for _, a := range albums {
		if a.ID == id {
			return a, nil
		}
	}
	return models.Album{}, errors.New("Album not found")
}
