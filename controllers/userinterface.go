package controllers

import "github.com/labstack/echo"

type InterfaceUsers interface {
	Create(c echo.Context) error
	Login(c echo.Context) error
}

