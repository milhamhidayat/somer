package router

import (
	controller "somer/interface/controller"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// NewRouter initalize router
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/offices", func(context echo.Context) error {
		return c.Office.Fetch(context)
	})

	return e
}
