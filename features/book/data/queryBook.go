package data

import (
	"alta/project2/features/book"
	"errors"

	"gorm.io/gorm"
)

type bookData struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) book.DataInterface {
	return &bookData{
		DB: conn,
	}
}

func (repo *bookData) SelectAll() ([]book.BookCore, error) {

	var dataBook []Book
	tx := repo.DB.Find(&dataBook)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataBookCore := toBookList(dataBook)

	return dataBookCore, nil

}

func (repo *bookData) SelectById(param int) (book.BookCore, error) {

	var data Book
	tx := repo.DB.First(&data, param)
	if tx.Error != nil {
		return book.BookCore{}, tx.Error
	}

	bookId := data.toBookUser()
	return bookId, nil

}

func (repo *bookData) CreateData(data book.BookCore, token int) (int, error) {

	var datacheck User
	txcheck := repo.DB.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return -1, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return -1, errors.New("not have access")
	}

	dataModel := InsBook(data)
	dataModel.UserID = uint(token)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *bookData) UpdateData(param, token int, dataUpdate book.BookCore) (int, error) {

	var datacheck Book
	tx := repo.DB.First(&datacheck, param)
	if tx.Error != nil {
		return -1, tx.Error
	}

	bookId := datacheck.toBookUser()

	if bookId.UserID == uint(token) {
		var data Book
		data.Title = dataUpdate.Title
		data.Author = dataUpdate.Author
		data.Publisher = dataUpdate.Publisher
		data.Page = dataUpdate.Page

		var book Book
		book.ID = dataUpdate.ID
		txUpdateId := repo.DB.Model(&book).Updates(data)
		if txUpdateId.Error != nil {
			return -1, txUpdateId.Error
		}

		var err error

		return int(txUpdateId.RowsAffected), err
	} else {
		return -1, errors.New("not have access")
	}

}

func (repo *bookData) DelData(param, token int) (int, error) {

	var datacheck Book
	tx := repo.DB.First(&datacheck, param)
	if tx.Error != nil {
		return -1, tx.Error
	}

	bookId := datacheck.toBookUser()

	if bookId.UserID == uint(token) {
		var data Book
		txDelId := repo.DB.Delete(&data, param)
		if txDelId.Error != nil {
			return -1, txDelId.Error
		}

		var err error

		return int(txDelId.RowsAffected), err
	} else {
		return -1, errors.New("not access")
	}

}
