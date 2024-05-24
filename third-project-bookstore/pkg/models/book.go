package models

import (
	"log"
	"time"

	"github.com/NeVajnoKak/Beginner-GoLang/tree/main/third-project-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Book struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name"`
	Author      string     `json:"author"`
	Publication string     `json:"publication"`
}

func init() {
	var err error
	config.Connect()
	db = config.GetDb()
	if db == nil {
		log.Fatalf("Failed to connect to the database.")
	}
	if err = db.AutoMigrate(&Book{}).Error; err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var book Book
	db := db.Where("id = ?", Id).Find(&book)
	if db.Error != nil {
		return nil, nil, db.Error
	}
	return &book, db, nil
}

func DeleteBook(ID int64) error {
	if err := db.Where("id = ?", ID).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}
