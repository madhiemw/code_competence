package controller

import (
	"codecompetence/middleware"
	"codecompetence/model/payload"
	"codecompetence/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func AddItem(c echo.Context) error {
	if _, err := middleware.IsUser(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	req := payload.CreateItemRequest{}

	c.Bind(&req)

	err := usecase.CreateItem(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to add item ")
	}
	resp := req
	return c.JSON(http.StatusOK, resp)
}

func GetItemByid(c echo.Context) error {
	id := c.Param("id")

	response, err := usecase.GetItemByid(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "item not found")
	}

	return c.JSON(http.StatusOK, response)
}
func GetItemByCategoryID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	response, err := usecase.GetItemByCategoryID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "item not found")
	}

	return c.JSON(http.StatusOK, response)
}
func GetItemByName(c echo.Context) error {
	name := c.QueryParam("keyword")

	response, err := usecase.GetItemByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "item not found")
	}

	return c.JSON(http.StatusOK, response)

}

func DeleteItem(c echo.Context) error {
	if _, err := middleware.IsUser(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}
	id := c.Param("id")

	item, err := usecase.GetItemByid(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "item not found")
	}

	err = usecase.DeleteItem(item)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "delete complete")
}

func GetAllItem(c echo.Context) error {
	response, err := usecase.GetAllItem()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateItem(c echo.Context) error {
	if _, err := middleware.IsUser(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}
	id := c.Param("id")

	item, err := usecase.GetItemByid(id)
	if err != nil {
		return err
	}
	c.Bind(item)
	response, err := usecase.UpdateItem(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
