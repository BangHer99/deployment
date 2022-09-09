package deliv

import (
	"alta/project2/features/user"
	"alta/project2/middlewares"
	"alta/project2/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	data user.ServiceInterface
}

func New(e *echo.Echo, usecase user.ServiceInterface) {

	handler := UserHandler{
		data: usecase,
	}

	e.GET("/users", handler.GetAllWithJWT, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetByIdWithJWT, middlewares.JWTMiddleware())
	e.POST("/users", handler.AddUser)
	e.PUT("/users/:id", handler.PutDataWithJWT, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeldateWithJWT, middlewares.JWTMiddleware())

}

func (users *UserHandler) GetAllWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	res, err := users.data.GetAll(idToken)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper(err))
	}

	respon := toResponList(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get all data", respon))

}

func (users *UserHandler) GetByIdWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	res, err := users.data.GetById(id, idToken)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	respon := toResponId(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))

}

func (users *UserHandler) AddUser(e echo.Context) error {

	var req UserReq
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper(err))
	}

	add := ToCore(req)
	row, errPost := users.data.PostData(add)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes insert data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper(errPost))
	}

}

func (users *UserHandler) PutDataWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	var req UserReq
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper(err))
	}

	var add user.UserCore
	if req.Email != "" {
		add.Email = req.Email
	}
	if req.Name != "" {
		add.Name = req.Name
	}
	if req.Password != "" {
		add.Password = req.Password
	}

	add.ID = uint(id)

	row, _ := users.data.PutData(id, idToken, add)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes update data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}

func (users *UserHandler) DeldateWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, _ := users.data.DeleteData(id, idToken)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes delete data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}
