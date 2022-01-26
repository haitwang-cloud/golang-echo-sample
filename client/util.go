package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func IntContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// RespError Convert error response into error message
func RespError(resp *resty.Response) error {
	httpError := &HttpError{
		Code:    resp.StatusCode(),
		Message: string(resp.Body()),
	}
	return fmt.Errorf("request failed [%d]: %s", httpError.Code, httpError.Message)
}
