package data

import (
	"alta/project2/features/user"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Email     string `gorm:"primary key"`
	ID        uint   `gorm:"autoIncrement"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Books     []Book `gorm:"foreignkey:UserID;references:ID"`
}

type Book struct {
	gorm.Model
	UserID    uint
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Page      int    `json:"page" form:"page"`
}

func ModelInsert(data user.UserCore) User {

	userData := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	return userData

}

func (dataBookUser *Book) toBookUser() user.BookCore {

	dataUser := user.BookCore{
		ID:        dataBookUser.ID,
		UserID:    dataBookUser.UserID,
		Title:     dataBookUser.Title,
		Author:    dataBookUser.Author,
		Publisher: dataBookUser.Publisher,
		Page:      dataBookUser.Page,
		CreatedAt: dataBookUser.CreatedAt,
		UpdatedAt: dataBookUser.UpdatedAt,
		DeletedAt: dataBookUser.DeletedAt.Time,
	}

	return dataUser

}

func toBookUserList(data []Book) []user.BookCore {
	var dataCore []user.BookCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toBookUser())
	}
	return dataCore
}

func (data *User) toUserCore(book []Book) user.UserCore {

	dataUser := user.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	bookUser := toBookUserList(book)

	for i, v := range bookUser {
		if v.UserID == dataUser.ID {
			dataUser.Book = append(dataUser.Book, bookUser[i])
		}
	}

	return dataUser
}

func toUserCoreList(data []User, book []Book) []user.UserCore {
	var dataCore []user.UserCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toUserCore(book))
	}
	return dataCore
}
