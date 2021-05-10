package router

import (
	"github.com/exuan/waka-api/api/handler"
	"github.com/exuan/waka-api/api/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &middleware.Context{Context: c}
			return next(cc)
		}
	})

	e.HTTPErrorHandler = handler.Error
	e.Any("/", handler.Welcome)
	e.Any("/api/heartbeat", handler.ApiHeartbeat)

	return e
}
