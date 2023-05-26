package usecase

import (
	"codecompetence/middleware"
	"codecompetence/model"
	"codecompetence/model/payload"
	"codecompetence/repository"
	"errors"
	"net/http"

	// "errors"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req payload.UserRegisterRequst) (resp payload.UserRegisterResponse, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	// if repository.EmailDuplicateChecker(req.Email) {
	// 	return resp, errors.New("email is already registered")
	// }
	newUser := &model.User{
		Email:    req.Email,
		Password: string(passwordHash),
	}

	err = repository.CreateUser(newUser)
	if err != nil {
		return
	}

	resp = payload.UserRegisterResponse{
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	return
}

func UserLogin(req payload.UserLoginRequest) (res payload.LoginUserResponse, err error) {
	user, err := repository.GetuserByEmail(req.Email)
	if err != nil {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Failed to generate token")
	}

	user.Token = token

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}
