package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authData "alta/project2/features/author/data"
	authDelivery "alta/project2/features/author/deliv"
	authService "alta/project2/features/author/usecase"

	bookData "alta/project2/features/book/data"
	bookDelivery "alta/project2/features/book/deliv"
	bookService "alta/project2/features/book/usecase"
	userData "alta/project2/features/user/data"
	userDelivery "alta/project2/features/user/deliv"
	userService "alta/project2/features/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userService.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	authDataFactory := authData.New(db)
	authUsecaseFactory := authService.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	bookDataFactory := bookData.New(db)
	bookUsecaseFactory := bookService.New(bookDataFactory)
	bookDelivery.New(e, bookUsecaseFactory)

}
