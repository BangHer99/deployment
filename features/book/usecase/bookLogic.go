package service

import (
	"alta/project2/features/book"
	"errors"
)

type bookService struct {
	dataBook book.DataInterface
}

func New(data book.DataInterface) book.ServiceInterface {
	return &bookService{
		dataBook: data,
	}

}

func (service *bookService) GetAll() ([]book.BookCore, error) {

	dataAll, err := service.dataBook.SelectAll()
	if err != nil {
		return nil, errors.New("failed get all data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data is still empty")
	} else {
		return dataAll, nil
	}

}

func (service *bookService) GetById(param int) (book.BookCore, error) {

	dataId, err := service.dataBook.SelectById(param)
	if err != nil {
		return book.BookCore{}, err
	}

	return dataId, nil

}

func (service *bookService) PostData(data book.BookCore, token int) (int, error) {

	if data.Title != "" && data.Author != "" && data.Publisher != "" && data.Page != 0 {

		add, err := service.dataBook.CreateData(data, token)
		if err != nil || add == 0 {
			return -1, err
		} else {
			return 1, nil
		}
	} else {
		return -1, errors.New("all input data must be filled")
	}

}

func (service *bookService) PutData(param, token int, data book.BookCore) (int, error) {

	row, err := service.dataBook.UpdateData(param, token, data)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil

}

func (service *bookService) DeleteData(param, token int) (int, error) {

	row, err := service.dataBook.DelData(param, token)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil

}
