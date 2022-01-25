package utils

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

type DropboxError struct {
	Error string
}
type AuthSuccess struct {
	/* variables */
}
type AuthError struct {
	/* variables */
}
type Article struct {
	Title   string
	Content string
	Author  string
	Tags    []string
}
type Error struct {
	/* variables */
}

//
// Package Level examples
//

func TestExampleGet(t *testing.T) {
	// Create a resty client
	client := resty.New()

	resp, err := client.R().Get("http://httpbin.org/get")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
}

func TestExampleEnhancedGet(t *testing.T) {
	// Create a resty client
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")

	printOutput(resp, err)
}

func Example_post() {
	// Create a resty client
	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		SetResult(AuthSuccess{}). // or SetResult(&AuthSuccess{}).
		Post("https://myapp.com/login")

	printOutput(resp, err)

	// POST []byte array
	// No need to set content type, if you have client level setting
	resp1, err1 := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
		SetResult(AuthSuccess{}). // or SetResult(&AuthSuccess{}).
		Post("https://myapp.com/login")

	printOutput(resp1, err1)

	// POST Struct, default is JSON content type. No need to set one
	resp2, err2 := client.R().
		SetBody(resty.User{Username: "testuser", Password: "testpass"}).
		SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).    // or SetError(AuthError{}).
		Post("https://myapp.com/login")

	printOutput(resp2, err2)

	// POST Map, default is JSON content type. No need to set one
	resp3, err3 := client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).    // or SetError(AuthError{}).
		Post("https://myapp.com/login")

	printOutput(resp3, err3)
}

func Example_put() {
	// Create a resty client
	client := resty.New()

	// Just one sample of PUT, refer POST for more combination
	// request goes as JSON content type
	// No need to set auth token, error, if you have client level settings
	resp, err := client.R().
		SetBody(Article{
			Title:   "go-resty",
			Content: "This is my article content, oh ya!",
			Author:  "Jeevanandam M",
			Tags:    []string{"article", "sample", "resty"},
		}).
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		SetError(&Error{}). // or SetError(Error{}).
		Put("https://myapp.com/article/1234")

	printOutput(resp, err)
}

func printOutput(resp *resty.Response, err error) {
	fmt.Println(resp, err)
}
