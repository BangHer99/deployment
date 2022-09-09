package service

import (
	"alta/project2/features/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	dataUser user.DataInterface
}

func New(data user.DataInterface) user.ServiceInterface {
	return &userService{
		dataUser: data,
	}

}

func (service *userService) GetAll(token int) ([]user.UserCore, error) {

	dataAll, err := service.dataUser.SelectAll(token)
	if err != nil {
		return nil, errors.New("failed get all data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data is still empty")
	} else {
		return dataAll, nil
	}

}

func (service *userService) GetById(param, token int) (user.UserCore, error) {

	dataId, err := service.dataUser.SelectById(param, token)
	if err != nil {
		return user.UserCore{}, err
	}

	return dataId, nil

}

func (service *userService) PostData(data user.UserCore) (int, error) {

	if data.Email != "" && data.Name != "" && data.Password != "" {
		passByte := []byte(data.Password)
		hashPass, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
		data.Password = string(hashPass)
		add, err := service.dataUser.CreateData(data)
		if err != nil || add == 0 {
			return -1, err
		} else {
			return 1, nil
		}
	} else {
		return -1, errors.New("all input data must be filled")
	}

}

func (service *userService) PutData(param, token int, data user.UserCore) (int, error) {

	if param != token {
		return -1, errors.New("not have access")
	}

	row, err := service.dataUser.UpdateData(param, data)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil

}

func (service *userService) DeleteData(param, token int) (int, error) {

	if param != token {
		return -1, errors.New("not have access")
	}

	_, err := service.dataUser.DelData(param)
	if err != nil {
		return -1, err
	}

	return 1, nil

}
