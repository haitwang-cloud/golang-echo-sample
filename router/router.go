package router

import (
	"github.com/haitwang-cloud/golang-echo-sample/utils/middlewares"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, wrapper middlewares.Wrapper) {
	setBibleController(e, wrapper)
}

func setBibleController(e *echo.Echo, wrapper middlewares.Wrapper) {
	bible := NewBibleController(wrapper)
	e.GET(BibleResults, func(c echo.Context) error { return bible.GetResult(c) })
}
