package controller

import (
	"codecompetence/model/payload"
	"codecompetence/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func UserRegister(c echo.Context) error {
	req := payload.UserRegisterRequst{}

	c.Bind(&req)
	
	repsonse, err := usecase.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to create User")
	}
	return c.JSON(http.StatusOK, repsonse)
}

func UserLogin(c echo.Context) error {
	req := payload.UserLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	res, err := usecase.UserLogin(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to login")
	}

	return c.JSON(http.StatusOK, res)
}
