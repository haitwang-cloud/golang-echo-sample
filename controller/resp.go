package controller

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// RespError Convert error response into error message
func RespError(resp *resty.Response) error {
	httpError := &HttpError{
		Code:    resp.StatusCode(),
		Message: string(resp.Body()),
	}
	return fmt.Errorf("request failed [%d]: %s", httpError.Code, httpError.Message)
}
