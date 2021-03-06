package controller

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/haitwang-cloud/golang-echo-sample/utils"
	"github.com/haitwang-cloud/golang-echo-sample/utils/middlewares"
	"net/http"
	"time"
)

const (
	// DefaultHost is the default host for the client
	DefaultHost       = "https://bible-api.com/"
	MaxRetryCount     = 3
	RetryWaitDuration = 5 * time.Second
)

var ValidStatus = []int{http.StatusOK, http.StatusAccepted}

type BibleClient struct {
	client  *resty.Client
	wrapper middlewares.Wrapper
}

func NewBibleClient(wrapper middlewares.Wrapper) *BibleClient {
	client := resty.New()
	client.SetRetryCount(MaxRetryCount).
		SetRetryWaitTime(RetryWaitDuration).
		SetBaseURL(DefaultHost)
	client.AddRetryCondition(
		// RetryConditionFunc type is for retry condition function
		// input: non-nil Response OR request execution error
		func(r *resty.Response, err error) bool {
			return !utils.IntContains(ValidStatus, r.StatusCode())
		},
	)
	client.SetError(&HttpError{})

	return &BibleClient{client: client, wrapper: wrapper}
}

func (c *BibleClient) GetResult(book, chapter, verse string) (*BiBleResult, error) {
	resp, err := c.client.R().SetResult(&BiBleResult{}).Get("/" + book + chapter + ":" + verse)
	if err != nil {
		return nil, fmt.Errorf("get failed: %s", err)
	}
	if resp.Error() != nil {
		c.wrapper.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, RespError(resp)
	}
	return resp.Result().(*BiBleResult), nil

}

func (c *BibleClient) PutResult(result BiBleResult) (*BiBleResult, error) {
	resp, err := c.client.R().SetBody(map[string]interface{}{
		"username": "jeeva@myjeeva.com",
		"address":  result,
	}).SetResult(&BiBleResult{}).Put("/")
	if err != nil {
		c.wrapper.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, fmt.Errorf("get failed: %s", err)

	}
	if resp.Error() != nil {
		c.wrapper.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, RespError(resp)
	}
	return resp.Result().(*BiBleResult), nil

}
