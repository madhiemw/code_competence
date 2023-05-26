package routes

import (
	"codecompetence/controller"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.POST("/add", controller.AddItem)
	e.GET("/items/:id", controller.GetItemByid)
	e.GET("/items",controller.GetAllItem)
	e.PUT("/items/:id", controller.UpdateItem)
	e.DELETE("/items/:id", controller.DeleteItem)
	e.POST("/register", controller.UserRegister)
	e.GET("/login", controller.UserLogin)
}
