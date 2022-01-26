package client

import (
	"fmt"
	"testing"
)

func TestBibleClient(t *testing.T) {
	// Create a resty client
	book, chapter, verse := "John", "3", "100000000"
	client := NewBibleClient()
	getResult, err := client.GetResult(book, chapter, verse)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Println(getResult)
}
