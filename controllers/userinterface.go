package controllers

import "github.com/labstack/echo"

type InterfaceUsers interface {
	CreateUser(c echo.Context) error
}
