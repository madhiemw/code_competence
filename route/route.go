package routes

import (
	"codecompetence/controller"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.POST("/register", controller.UserRegister)
	e.POST("/add", controller.AddItem)
	e.GET("/items/:id", controller.GetItemByid)
	e.GET("/items",controller.GetAllItem)
	e.GET("/items", controller.GetItemByName)
	e.GET("/login", controller.UserLogin)
	e.PUT("/items/:id", controller.UpdateItem)
	e.DELETE("/del/items/:id", controller.DeleteItem)
}
