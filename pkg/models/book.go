package models

import (
	"errors"

	"github.com/d28035203/scaling-waddle/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book is a bookstore catalog entry.
type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:255;not null"`
	Author      string `json:"author" gorm:"size:255;not null"`
	Publication string `json:"publication" gorm:"size:255"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	_ = db.AutoMigrate(&Book{})
}

// CreateBook inserts a new book row.
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks returns every book in the catalog.
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById loads a single book by primary key.
func GetBookById(id int64) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, result.Error
	}
	return &book, nil
}

// UpdateBook persists changes to an existing book.
func UpdateBook(b *Book) error {
	return db.Save(b).Error
}

// DeleteBook removes a book by id and returns the deleted record.
func DeleteBook(id int64) Book {
	var book Book
	db.First(&book, id)
	db.Delete(&book)
	return book
}
