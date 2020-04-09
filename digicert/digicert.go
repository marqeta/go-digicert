package digicert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://www.digicert.com/services/v2/"
	headerAPIKey   = "X-DC-DEVKEY"
)

type DigicertClient interface {
	NewRequest(string, string, interface{}) (*http.Request, error)
	Do(*http.Request, interface{}) (*Response, error)
}

type Client struct {
	httpClient HTTPClient
	BaseURL    *url.URL

	common service

	apiKey string

	Users  *UsersService
	Orders *OrdersService
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type service struct {
	client DigicertClient
}

func NewClient(apiKey string, httpClient HTTPClient) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{httpClient: httpClient, BaseURL: baseURL}
	c.apiKey = apiKey
	c.common.client = c
	c.Users = (*UsersService)(&c.common)
	c.Orders = (*OrdersService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set(headerAPIKey, c.apiKey)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{Response: resp}
	if v != nil {
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return response, err
}

type Response struct {
	*http.Response
}
