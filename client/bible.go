package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-web-sample/utils"
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
	client *resty.Client
}

func NewBibleClient() *BibleClient {
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

	return &BibleClient{client: client}
}

func (c *BibleClient) GetResult(book, chapter, verse string) (*Result, error) {
	resp, err := c.client.R().SetResult(&Result{}).Get("/" + book + chapter + ":" + verse)
	if err != nil {
		return nil, fmt.Errorf("get failed: %s", err)
	}
	if resp.Error() != nil {
		return nil, RespError(resp)
	}
	return resp.Result().(*Result), nil

}

func (c *BibleClient) PutResult(result Result) (*Result, error) {
	resp, err := c.client.R().SetBody(map[string]interface{}{
		"username": "jeeva@myjeeva.com",
		"address":  result,
	}).SetResult(&Result{}).Put("/")
	if err != nil {
		return nil, fmt.Errorf("get failed: %s", err)
	}
	if resp.Error() != nil {
		return nil, RespError(resp)
	}
	return resp.Result().(*Result), nil

}
