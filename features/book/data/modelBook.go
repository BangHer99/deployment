package data

import (
	// "alta/project2/features/book"
	"alta/project2/features/book"
	_ "alta/project2/features/book"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	UserID    uint
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Page      int    `json:"page" form:"page"`
}

type User struct {
	Email     string `gorm:"primary key"`
	ID        uint   `gorm:"autoIncrement"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func InsBook(data book.BookCore) Book {
	userData := Book{
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		Page:      data.Page,
	}

	return userData
}

func (dataBook *Book) toBookUser() book.BookCore {

	dataBookCore := book.BookCore{
		ID:        dataBook.ID,
		UserID:    dataBook.UserID,
		Title:     dataBook.Title,
		Author:    dataBook.Author,
		Publisher: dataBook.Publisher,
		Page:      dataBook.Page,
		CreatedAt: dataBook.CreatedAt,
		UpdatedAt: dataBook.UpdatedAt,
		DeletedAt: dataBook.DeletedAt.Time,
	}

	return dataBookCore

}

func toBookList(data []Book) []book.BookCore {
	var dataCore []book.BookCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toBookUser())
	}
	return dataCore
}
