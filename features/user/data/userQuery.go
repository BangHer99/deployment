package data

import (
	"alta/project2/features/user"
	"errors"

	"gorm.io/gorm"
)

type userData struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) user.DataInterface {
	return &userData{
		DB: conn,
	}
}

func (repo *userData) book() []Book {

	var dataBookUser []Book
	tx := repo.DB.Find(&dataBookUser)
	if tx.Error != nil {
		return nil
	}

	return dataBookUser

}

func (repo *userData) SelectAll(token int) ([]user.UserCore, error) {

	var datacheck User
	txcheck := repo.DB.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return nil, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return nil, errors.New("not have access")
	}

	var dataAll []User
	tx := repo.DB.Find(&dataAll)
	if tx.Error != nil {
		return nil, tx.Error
	}

	bookList := repo.book()

	dataCore := toUserCoreList(dataAll, bookList)
	return dataCore, nil

}

func (repo *userData) SelectById(param, token int) (user.UserCore, error) {

	var datacheck User
	txcheck := repo.DB.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return user.UserCore{}, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return user.UserCore{}, errors.New("not have access")
	}

	var data User
	tx := repo.DB.First(&data, param)
	if tx.Error != nil {
		return user.UserCore{}, tx.Error
	}

	bookList := repo.book()

	userId := data.toUserCore(bookList)
	return userId, nil

}

func (repo *userData) CreateData(data user.UserCore) (int, error) {

	dataModel := ModelInsert(data)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *userData) UpdateData(param int, dataUpdate user.UserCore) (int, error) {

	var data User
	data.Name = dataUpdate.Name
	data.Email = dataUpdate.Email
	data.Password = dataUpdate.Password

	var user User
	user.ID = dataUpdate.ID
	txUpdateId := repo.DB.Model(&user).Updates(data)
	if txUpdateId.Error != nil {
		return -1, txUpdateId.Error
	}

	var err error

	return int(txUpdateId.RowsAffected), err

}

func (repo *userData) DelData(param int) (int, error) {

	var data User
	txDelId := repo.DB.Delete(&data, param)
	if txDelId.Error != nil {
		return -1, txDelId.Error
	}

	var book Book
	tx := repo.DB.Where("user_id=?", param).Delete(&book)
	if tx.Error != nil {
		return -1, tx.Error
	}

	return int(txDelId.RowsAffected), nil

}
