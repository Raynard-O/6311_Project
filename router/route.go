package router

import "github.com/labstack/echo"

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.JSON(200, "welcome home")
	})

	return e
}
