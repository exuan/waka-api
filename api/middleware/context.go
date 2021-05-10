package middleware

import (
	"net/http"

	"github.com/exuan/waka-api/api/message"
	"github.com/exuan/waka-api/api/response"
	"github.com/gin-gonic/gin/binding"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type (
	Context struct {
		echo.Context
	}
)

func (c *Context) filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func (c *Context) Bind(i interface{}) error {
	b := binding.Default(c.Request().Method, c.filterFlags(c.Request().Header.Get(echo.HeaderContentType)))
	if err := b.Bind(c.Request(), i); err != nil {
		return err
	}
	return binding.Validator.ValidateStruct(i)
	//return c.Validate(i)
}

func (c *Context) JSON(code int, d interface{}) error {
	if d == nil {
		d = `{}`
	}

	res := &response.Res{
		Error: code,
		Data:  d,
	}

	status := http.StatusOK
	//@todo refactor
	if res.Msg = message.HttpText(code); code != message.OK && len(res.Msg) == 0 {
		status = code
		res.Msg = http.StatusText(code)
	}

	bs, _ := jsoniter.Marshal(res)
	return c.JSONBlob(status, bs)
}
