package controller

import (
	"codecompetence/model/payload"
	"codecompetence/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func AddItem(c echo.Context) error {
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
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	_, err := usecase.GetItemByid(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "item not found")
	}
	return nil
}

func DeleteItem(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

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
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

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
