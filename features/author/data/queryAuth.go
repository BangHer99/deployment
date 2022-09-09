package data

import (
	"alta/project2/features/author"

	"log"

	"gorm.io/gorm"
)

type userLogin struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) author.DataInterface {

	return &userLogin{
		DB: conn,
	}

}

func (repo *userLogin) LoginUser(email, password string) (author.AuthCore, error) {

	var auth User
	txEmail := repo.DB.Where("email = ?", email).First(&auth)
	if txEmail.Error != nil {
		log.Println("Error tx")
		return author.AuthCore{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		log.Println("Error row")
		return author.AuthCore{}, txEmail.Error
	}

	var data = toCore(auth)

	return data, nil

}
