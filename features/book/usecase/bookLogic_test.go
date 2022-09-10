package service

import (
	"alta/project2/features/book"
	"alta/project2/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	bookMock := new(mocks.BookData)
	returnData := []book.BookCore{{Title: "Bumi", Author: "tere liye", Publisher: "gramedia", Page: 280}}

	t.Run("Get All Success", func(t *testing.T) {

		bookMock.On("SelectAll").Return(returnData, nil).Once()

		useCase := New(bookMock)
		res, err := useCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		bookMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		bookMock.On("SelectAll").Return(nil, nil).Once()

		useCase := New(bookMock)
		res, _ := useCase.GetAll()
		assert.Equal(t, 0, len(res))
		bookMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		bookMock.On("SelectAll").Return(nil, errors.New("error")).Once()

		useCase := New(bookMock)
		_, err := useCase.GetAll()
		assert.Error(t, err)
		bookMock.AssertExpectations(t)

	})

}

func TestGetById(t *testing.T) {

	bookMock := new(mocks.BookData)
	returnData := book.BookCore{ID: 1, Title: "Bumi", Author: "tere liye", Publisher: "gramedia", Page: 280}
	param := 1

	t.Run("Get by id success", func(t *testing.T) {
		bookMock.On("SelectById", param).Return(returnData, nil).Once()

		useCase := New(bookMock)
		res, _ := useCase.GetById(param)
		assert.Equal(t, param, int(res.ID))
		bookMock.AssertExpectations(t)

	})

	t.Run("Get by id failed", func(t *testing.T) {

		bookMock.On("SelectById", param).Return(book.BookCore{}, errors.New("error")).Once()

		useCase := New(bookMock)
		param := 1
		res, err := useCase.GetById(param)
		assert.Error(t, err)
		assert.NotEqual(t, param, int(res.ID))
		bookMock.AssertExpectations(t)

	})

}

func TestPut(t *testing.T) {

	bookMock := new(mocks.BookData)
	input := book.BookCore{Title: "Bumi", Author: "tere liye", Publisher: "gramedia", Page: 280}
	param := 1
	token := 1

	t.Run("update succes", func(t *testing.T) {

		bookMock.On("UpdateData", param, token, input).Return(1, nil).Once()

		useCase := New(bookMock)
		res, _ := useCase.PutData(param, token, input)
		assert.Equal(t, 1, res)
		bookMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		bookMock.On("UpdateData", param, token, input).Return(-1, errors.New("error")).Once()

		useCase := New(bookMock)
		res, err := useCase.PutData(param, token, input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		bookMock.AssertExpectations(t)

	})

}

func TestPostData(t *testing.T) {

	bookMock := new(mocks.BookData)
	input := book.BookCore{Title: "Bumi", Author: "tere liye", Publisher: "gramedia", Page: 280}
	token := 1

	t.Run("create success", func(t *testing.T) {

		bookMock.On("CreateData", input, token).Return(1, nil).Once()

		useCase := New(bookMock)
		res, _ := useCase.PostData(input, token)
		assert.Equal(t, 1, res)
		bookMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		bookMock.On("CreateData", input, token).Return(-1, errors.New("error")).Once()

		useCase := New(bookMock)
		res, err := useCase.PostData(input, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		bookMock.AssertExpectations(t)

	})

	t.Run("create failed", func(t *testing.T) {

		input.Title = ""
		input.Author = ""
		useCase := New(bookMock)
		res, err := useCase.PostData(input, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		bookMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	bookMock := new(mocks.BookData)
	token := 1
	param := 3

	t.Run("delete succes", func(t *testing.T) {

		bookMock.On("DelData", param, token).Return(1, nil).Once()

		useCase := New(bookMock)
		res, _ := useCase.DeleteData(param, token)
		assert.Equal(t, 1, res)
		bookMock.AssertExpectations(t)

	})

	t.Run("delete succes", func(t *testing.T) {

		bookMock.On("DelData", param, token).Return(-1, errors.New("error")).Once()

		useCase := New(bookMock)
		res, err := useCase.DeleteData(param, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		bookMock.AssertExpectations(t)

	})

}
