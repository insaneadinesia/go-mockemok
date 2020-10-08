package handler

import (
	"github.com/insaneadinesia/go-mockemok/config"
	"github.com/insaneadinesia/go-mockemok/utils"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	request := c.Get("Request")
	req := request.(config.MockGroupRequest)
	resp := utils.GetResponseBody(req.Body)

	return c.JSON(req.Status, resp)
}

func Post(c echo.Context) error {
	request := c.Get("Request")
	req := request.(config.MockGroupRequest)
	resp := utils.GetResponseBody(req.Body)

	return c.JSON(req.Status, resp)
}

func Put(c echo.Context) error {
	request := c.Get("Request")
	req := request.(config.MockGroupRequest)
	resp := utils.GetResponseBody(req.Body)

	return c.JSON(req.Status, resp)
}

func Patch(c echo.Context) error {
	request := c.Get("Request")
	req := request.(config.MockGroupRequest)
	resp := utils.GetResponseBody(req.Body)

	return c.JSON(req.Status, resp)
}

func Delete(c echo.Context) error {
	request := c.Get("Request")
	req := request.(config.MockGroupRequest)
	resp := utils.GetResponseBody(req.Body)

	return c.JSON(req.Status, resp)
}
