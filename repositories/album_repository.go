package repositories

import "github.com/denisemignoli/restful-api-with-go-and-gin/models"

type AlbumRepository interface {
	GetAlbums() []models.Album
	SaveAlbums(newAlbum models.Album) (int64, error)
	GetAlbumByID(id int64) (*models.Album, error)
}
