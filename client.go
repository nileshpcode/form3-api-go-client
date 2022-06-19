package main

import (
	"github.com/go-resty/resty/v2"
)

// Client contains the interface of different form3 API
type Client struct {
	AccountInterface
}

// NewClient creates a new client set for the given http client
// baseURL is the root URL for all invocations of the client
func NewClient(baseURL string) Client {
	c := resty.New().SetBaseURL(baseURL)
	return Client{
		AccountInterface: newAccountService(c),
	}
}
