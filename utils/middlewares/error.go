package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type APIError struct {
	Code    int
	Message string
}

// ErrorController is a controller for handling errors.
type ErrorController struct {
	wrapper Wrapper
}

// NewErrorController is constructor.
func NewErrorController(wrapper Wrapper) *ErrorController {
	return &ErrorController{wrapper: wrapper}
}

// JSONError is cumstomize error handler
func (controller *ErrorController) JSONError(err error, c echo.Context) {
	logger := controller.wrapper.GetLogger()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}

	var apierr APIError
	apierr.Code = code
	apierr.Message = msg

	if !c.Response().Committed {
		if reserr := c.JSON(code, apierr); reserr != nil {
			logger.GetZapLogger().Errorf(reserr.Error())
		}
	}
	logger.GetZapLogger().Debugf(err.Error())
}
