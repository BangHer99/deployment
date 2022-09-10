package service

import (
	"alta/project2/features/user"
	"alta/project2/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {

	userMock := new(mocks.UserData)
	returnData := []user.UserCore{{Name: "Hery", ID: 1, Email: "Hery@gmail.com", Password: "asdf"}}
	token := 1
	t.Run("Get All Success", func(t *testing.T) {

		userMock.On("SelectAll", token).Return(returnData, nil).Once()

		useCase := New(userMock)
		res, err := useCase.GetAll(token)
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		userMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		userMock.On("SelectAll", token).Return(nil, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.GetAll(token)
		assert.Equal(t, 0, len(res))
		userMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		userMock.On("SelectAll", token).Return(nil, errors.New("error")).Once()

		useCase := New(userMock)
		_, err := useCase.GetAll(token)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestGetById(t *testing.T) {

	userMock := new(mocks.UserData)
	returnData := user.UserCore{Name: "Hery", ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
	param := 1
	token := 1

	t.Run("Get by id success", func(t *testing.T) {
		userMock.On("SelectById", param, token).Return(returnData, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.GetById(param, token)
		assert.Equal(t, param, int(res.ID))
		userMock.AssertExpectations(t)

	})

	t.Run("Get by id failed", func(t *testing.T) {

		userMock.On("SelectById", param, token).Return(user.UserCore{}, errors.New("error")).Once()

		useCase := New(userMock)
		param := 1
		res, err := useCase.GetById(param, token)
		assert.Error(t, err)
		assert.NotEqual(t, param, int(res.ID))
		userMock.AssertExpectations(t)

	})

}

func TestPut(t *testing.T) {

	userMock := new(mocks.UserData)
	input := user.UserCore{Name: "Hery", ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
	param := 1
	token := 1

	t.Run("update succes", func(t *testing.T) {

		userMock.On("UpdateData", param, input).Return(1, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.PutData(param, token, input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		userMock.On("UpdateData", param, input).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res, err := useCase.PutData(param, token, input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		token = 2
		useCase := New(userMock)
		res, err := useCase.PutData(param, token, input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestPostData(t *testing.T) {

	userMock := new(mocks.UserData)
	input := user.UserCore{Name: "Hery", Email: "Hery@gmail.com", Password: "asdf"}

	t.Run("create success", func(t *testing.T) {

		userMock.On("CreateData", mock.Anything).Return(1, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.PostData(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		userMock.On("CreateData", mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res, err := useCase.PostData(input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

	t.Run("create failed", func(t *testing.T) {

		input.Name = ""
		input.Password = ""
		useCase := New(userMock)
		res, err := useCase.PostData(input)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	userMock := new(mocks.UserData)
	token := 1
	param := 1

	t.Run("delete succes", func(t *testing.T) {

		userMock.On("DelData", param).Return(1, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.DeleteData(param, token)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		userMock.On("DelData", param).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res, err := useCase.DeleteData(param, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		token = 3
		useCase := New(userMock)
		res, err := useCase.DeleteData(param, token)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}
