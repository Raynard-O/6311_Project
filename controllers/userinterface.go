package controllers

import "github.com/labstack/echo"

type InterfaceUsers interface {
	Create(c echo.Context) error
	Search(c echo.Context) error
	Select(c echo.Context) error
}

