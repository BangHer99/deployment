package deliv

import (
	"alta/project2/features/author"
	"alta/project2/utils/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userUsecase author.UsecaseInterface
}

func New(e *echo.Echo, usecase author.UsecaseInterface) {

	handler := AuthHandler{
		userUsecase: usecase,
	}

	e.POST("/auth", handler.Auth)

}

func (h *AuthHandler) Auth(c echo.Context) error {

	var req Request
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, errBind)
	}

	str, err := h.userUsecase.LoginAuthorized(req.Email, req.Password)
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", str))

}
