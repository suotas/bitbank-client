package client

import (
	"errors"
	"net/url"
	"net/http"
)

const (
	baseUrl = "https://public.bitbank.cc"
	privateBaseUrl = "https://api.bitbank.cc/v1"
)

type Client struct {
	URL			*url.URL
	PrivateURL	*url.URL
	ApiToken	string
	ApiSecret	string
	HTTPClient	*http.Client
}

func NewClient(apiToken, apiSecret string) (*Client, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("apiToken is required")
	}
	publicUrl, _ := url.ParseRequestURI(baseUrl)

	if len(apiSecret) == 0 {
		return nil, errors.New("apiSecret is required")
	}
	privateUrl, _ := url.ParseRequestURI(privateBaseUrl)

	client := &http.Client{}

	return &Client{
		URL: publicUrl,
		PrivateURL: privateUrl,
		ApiToken: apiToken,
		ApiSecret: apiSecret,
		HTTPClient: client,
	},nil
}

func (c *Client) SendRequest(path string) (*http.Response, error){
	res, _ := c.HTTPClient.Get(path)
	return res, nil
}