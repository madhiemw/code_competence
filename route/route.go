package routes

import (
	"codecompetence/controller"
	"codecompetence/middleware"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.POST("/register", controller.UserRegister)
	e.GET("/items/:id", controller.GetItemByid)
	e.GET("/items",controller.GetAllItem)
	e.GET("/items", controller.GetItemByName)
	e.GET("/login", controller.UserLogin)
	e.GET("/category/:id ", controller.GetItemByCategoryID)

	
	jwt := e.Group("user", middleware.IsLoggedIn)
	jwt.POST("/add", controller.AddItem)
	jwt.DELETE("/del/items/:id", controller.DeleteItem)
	jwt.PUT("/items/:id", controller.UpdateItem)
}
