package service

import (
	"alta/project2/features/author"
	"alta/project2/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginAuthorized(t *testing.T) {

	dataLayerMock := new(mocks.AuthData)
	DataMock := author.AuthCore{ID: 1, Email: "Hery@gmail.com", Password: "$2a$12$laJ1gBshMH0dthChQQBmAudSWjSKbYSFK.Jm7HcjcKuA7R6Mxrwu6"}

	t.Run("LoginUserSuccess", func(t *testing.T) {
		dataLayerMock.On("LoginUser", mock.Anything, mock.Anything).Return(DataMock, nil).Once()

		Data := author.AuthCore{ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
		usecase := New(dataLayerMock)
		_, err := usecase.LoginAuthorized(Data.Email, Data.Password)
		assert.NoError(t, err)
		dataLayerMock.AssertExpectations(t)
	})

	t.Run("LoginUserFailed", func(t *testing.T) {

		Data := author.AuthCore{ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
		usecase := New(dataLayerMock)
		resultData, err := usecase.LoginAuthorized("", Data.Password)
		assert.NoError(t, err)
		assert.Equal(t, resultData, "")

		dataLayerMock.AssertExpectations(t)
	})

	t.Run("LoginUserFailed", func(t *testing.T) {

		Data := author.AuthCore{ID: 1, Email: "Hery@mail.com", Password: "asdf"}
		usecase := New(dataLayerMock)
		resultData, err := usecase.LoginAuthorized("", Data.Password)
		assert.NoError(t, err)
		assert.Equal(t, resultData, "")
		dataLayerMock.AssertExpectations(t)

	})

	t.Run("LoginUserFailed", func(t *testing.T) {
		dataLayerMock.On("LoginUser", mock.Anything, mock.Anything).Return(DataMock, errors.New("error")).Once()

		Data := author.AuthCore{ID: 1, Email: "Hery@mail.com", Password: "asdf"}
		usecase := New(dataLayerMock)
		resultData, err := usecase.LoginAuthorized(Data.Email, Data.Password)
		assert.Error(t, err)
		assert.Equal(t, resultData, "")
		dataLayerMock.AssertExpectations(t)

	})

	t.Run("LoginUserFailed", func(t *testing.T) {
		DataMock.Password = "1234124"
		dataLayerMock.On("LoginUser", mock.Anything, mock.Anything).Return(DataMock, nil).Once()

		Data := author.AuthCore{ID: 1, Email: "Hery@gmail.com", Password: "asdf"}
		usecase := New(dataLayerMock)
		resultData, _ := usecase.LoginAuthorized(Data.Email, Data.Password)
		assert.NotEqual(t, resultData, "Hery@gmail.com")
		dataLayerMock.AssertExpectations(t)

	})

}
