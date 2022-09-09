package deliv

import (
	"alta/project2/features/book"
	"alta/project2/middlewares"
	"alta/project2/utils/helper"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	data book.ServiceInterface
}

func New(e *echo.Echo, usecase book.ServiceInterface) {

	handler := BookHandler{
		data: usecase,
	}

	e.GET("/books", handler.GetAllWithJWT)
	e.GET("/books/:id", handler.GetByIdWithJWT)
	e.POST("/books", handler.AddUser, middlewares.JWTMiddleware())
	e.PUT("/books/:id", handler.PutDataWithJWT, middlewares.JWTMiddleware())
	e.DELETE("/books/:id", handler.DeldateWithJWT, middlewares.JWTMiddleware())

}

func (books *BookHandler) GetAllWithJWT(e echo.Context) error {

	res, err := books.data.GetAll()
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("tidak ada data"))
	}

	respon := toResponList(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get all data", respon))

}

func (books *BookHandler) GetByIdWithJWT(e echo.Context) error {

	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	res, err := books.data.GetById(id)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	respon := toRespon(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))

}

func (books *BookHandler) AddUser(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	var req Request
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper(err))
	}

	add := toCore(req)
	row, _ := books.data.PostData(add, idToken)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes insert data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("failed create book"))
	}

}

func (books *BookHandler) PutDataWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	var req Request
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	var add book.BookCore
	if req.Title != "" {
		add.Title = req.Title
	}
	if req.Author != "" {
		add.Author = req.Author
	}
	if req.Publisher != "" {
		add.Publisher = req.Publisher
	}
	if req.Page != 0 {
		add.Page = req.Page
	}

	add.ID = uint(id)

	row, _ := books.data.PutData(id, idToken, add)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes update data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}

func (books *BookHandler) DeldateWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, _ := books.data.DeleteData(id, idToken)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes delete data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}
