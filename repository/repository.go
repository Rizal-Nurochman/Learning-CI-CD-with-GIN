package repository

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string
	Writer        string
	YearPublished string
}

var db *gorm.DB

func ConnectDB() error {
	var err error

	db, err = gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Book{})
	if err != nil {
		return err
	}

	log.Println("Connect DB Berhasil!")
	return nil
}

func GetAll() ([]Book, error) {
	var books []Book

	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}