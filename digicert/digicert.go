package digicert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://www.digicert.com/services/v2/"
	headerAPIKey   = "X-DC-DEVKEY"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type DigicertClient interface {
	NewRequest(string, string, interface{}) (*http.Request, error)
	Do(*http.Request, interface{}) (*Response, error)
}

type Client struct {
	httpClient HTTPClient
	BaseURL    *url.URL

	common service

	apiKey string

	Users         *UsersService
	Orders        *OrdersService
	Organizations *OrganizationsService
	Products      *ProductsService
	Certificates  *CertificatesService
}

func NewClient(apiKey string, httpClient HTTPClient, apiURL string) (*Client, error) {
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{httpClient: httpClient, BaseURL: parsedURL}
	c.apiKey = apiKey
	c.common.client = c
	c.Users = (*UsersService)(&c.common)
	c.Orders = (*OrdersService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.Products = (*ProductsService)(&c.common)
	c.Certificates = (*CertificatesService)(&c.common)
	return c, nil
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
	response := &Response{Response: resp}

	err = c.CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		defer resp.Body.Close()
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

func (c *Client) CheckResponse(resp *http.Response) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		errorResponse.rawBody = data
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

type Response struct {
	*http.Response
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	*http.Response
	Errors  []APIError `json:"errors"`
	rawBody []byte
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %+v %d %v %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.Header, r.Response.StatusCode, r.Errors, r.rawBody)
}

type service struct {
	client DigicertClient
}

func executeAction(c DigicertClient, method, path string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}
