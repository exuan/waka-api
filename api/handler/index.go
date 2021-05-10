package handler

import (
	"fmt"
	"net/http"

	"github.com/exuan/waka-api/api"
	"github.com/exuan/waka-api/api/message"
	"github.com/labstack/echo/v4"
)

func Welcome(c echo.Context) error {
	return c.JSON(message.OK, api.Cfg.SrvEnv)
}

func Error(err error, c echo.Context) {
	if !c.Response().Committed {
		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		} else {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}

		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else {
			err = c.JSONBlob(he.Code, []byte(fmt.Sprintf(message.ResFormat, he.Code, he.Message, `{}`)))
		}
	}
}
