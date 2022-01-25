package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	// DefaultHost is the default host for the client
	DefaultHost       = "https://international.v1.hitokoto.cn"
	MaxRetryCount     = 3
	RetryWaitDuration = 5 * time.Second
)

var ValidStatus = []int{http.StatusOK, http.StatusAccepted}

type SampleCatClient struct {
	client *resty.Client
}

func NewSampleCatClient() *SampleCatClient {
	client := resty.New()
	client.SetRetryCount(MaxRetryCount).
		SetRetryWaitTime(RetryWaitDuration).
		SetBaseURL(DefaultHost)
	client.AddRetryCondition(
		// RetryConditionFunc type is for retry condition function
		// input: non-nil Response OR request execution error
		func(r *resty.Response, err error) bool {
			return !intContains(ValidStatus, r.StatusCode())
		},
	)
	client.SetError(&Error{})

	return &SampleCatClient{client: client}
}

func (c *SampleCatClient) GetResult() (*Result, error) {
	resp, err := c.client.R().SetResult(&Result{}).Get("/")
	if err != nil {
		return nil, fmt.Errorf("get failed: %s", err)
	}
	if resp.Error() != nil {
		return nil, APIError(resp)
	}
	return resp.Result().(*Result), nil

}

type Error struct {
	Code    string `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

func intContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Convert error response into error message
func APIError(resp *resty.Response) error {
	apiError := resp.Error().(*Error)
	return fmt.Errorf("request failed [%s]: %s", apiError.Code, apiError.Message)
}

type Result struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	Hitokoto   string      `json:"hitokoto"`
	Type       string      `json:"type"`
	From       string      `json:"from"`
	FromWho    interface{} `json:"from_who"`
	Creator    string      `json:"creator"`
	CreatorUID int         `json:"creator_uid"`
	Reviewer   int         `json:"reviewer"`
	CommitFrom string      `json:"commit_from"`
	CreatedAt  string      `json:"created_at"`
	Length     int         `json:"length"`
}
