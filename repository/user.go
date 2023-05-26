package repository

import (
	"codecompetence/config"
	"codecompetence/model"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(user *model.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetuserByEmail(email string) (user *model.User, err error) {
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func EmailDuplicateChecker(email string) bool {
	var count int64
	user := model.User{}
	if err := config.DB.Model(&user).Where("email = ?", email).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}
	return count == 0
}
