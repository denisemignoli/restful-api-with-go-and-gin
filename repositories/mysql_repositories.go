package repositories

import (
	"database/sql"
	"github.com/denisemignoli/restful-api-with-go-and-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type AlbumMySQLRepository struct {
	db *sql.DB
}

func NewAlbumMySQLRepository() *AlbumMySQLRepository {
	db, err := sql.Open("mysql", "root:code2022@tcp(localhost:3306)/bd_albums")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &AlbumMySQLRepository{
		db: db,
	}
}

func (a *AlbumMySQLRepository) GetAlbums() []models.Album {

	return albums
}

func (a *AlbumMySQLRepository) SaveAlbums(newAlbum models.Album) {
	albums = append(albums, newAlbum)
}

func (a *AlbumMySQLRepository) GetAlbumByID(id string) (models.Album, error) {
	return models.Album{}, nil
}
