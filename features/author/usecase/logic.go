package service

import (
	"alta/project2/features/author"
	"alta/project2/middlewares"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type authorizedUsecase struct {
	authorizedData author.DataInterface
}

func New(data author.DataInterface) author.UsecaseInterface {
	return &authorizedUsecase{
		authorizedData: data,
	}
}

func (usecase *authorizedUsecase) LoginAuthorized(email, password string) (string, error) {

	var err error
	if email == "" || password == "" {
		return "", err
	}

	results, errEmail := usecase.authorizedData.LoginUser(email, password)
	if errEmail != nil {
		return "", errEmail
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		log.Println("Error pw")
		return "", err
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	if errToken != nil {
		return "", err
	}

	if token == "" {
		return "", err
	}

	return token, nil

}
