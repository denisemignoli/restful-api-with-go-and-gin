package repositories

import (
	"database/sql"
	"fmt"
	"github.com/denisemignoli/restful-api-with-go-and-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
)

type AlbumMySQLRepository struct {
	db *sql.DB
}

const (
	username = "root"
	password = "code2022"
	host     = "localhost"
	port     = 3306
	database = "db_albums"
)

func NewAlbumMySQLRepository() *AlbumMySQLRepository {
	// build the DNS
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	// open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("impossible to create the connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// execute the script to create the table
	err = executeSQLScript(db, "scripts/create_table.sql")
	if err != nil {
		log.Fatalf("failed to execute the script: %s", err)
	}

	return &AlbumMySQLRepository{
		db: db,
	}
}

func executeSQLScript(db *sql.DB, filename string) error {
	// read the file
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// execute the script
	_, err = db.Exec(string(b))
	return err
}

func (a *AlbumMySQLRepository) GetAlbums() []models.Album {
	var albums []models.Album

	rows, err := a.db.Query("SELECT * FROM `albums`")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var album models.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			log.Fatal(err)
		}
		albums = append(albums, album)
	}
	return albums
}

func (a *AlbumMySQLRepository) SaveAlbums(newAlbum models.Album) (int64, error) {
	// insert the album in the database
	res, err := a.db.Exec("INSERT INTO `albums` (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		return 0, err
	}
	// get the id of the last inserted album
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AlbumMySQLRepository) GetAlbumByID(id int64) (*models.Album, error) {
	album := models.Album{}
	err := a.db.QueryRow("SELECT * FROM `albums` WHERE id = ?", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		return nil, err
	}
	return &album, nil
}
