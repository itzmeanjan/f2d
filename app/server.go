package app

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Run - Start `f2d` system
func Run() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")

	{

		v1.GET("/hello", func(c echo.Context) error {
			return nil
		})

	}

}
