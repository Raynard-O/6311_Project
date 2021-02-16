package router

import (
	"6311_Project/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	author := controllers.InterfaceUsers(controllers.NewAuthor())

	e.GET("/", author.CreateUser)

	return e
}
