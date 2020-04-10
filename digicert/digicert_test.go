package digicert

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) NewRequest(method string, path string, opt interface{}) (*http.Request, error) {
	ret := m.Called(method, path, opt)
	return ret.Get(0).(*http.Request), ret.Error(1)

}

func (m *MockClient) Do(req *http.Request, v interface{}) (*Response, error) {
	ret := m.Called(req, v)
	return ret.Get(0).(*Response), ret.Error(1)
}

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	ret := m.Called(req)
	return ret.Get(0).(*http.Response), ret.Error(1)
}

func TestNewClient(t *testing.T) {
	c := NewClient("", nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	c2 := NewClient("", nil)
	if c.httpClient == c2.httpClient {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient("secret123", nil)
	username := "u"

	inURL, outURL := "foo", defaultBaseURL+"foo"
	inBody, outBody := &User{Username: username}, `{"username":"u"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test that api key is attached to the request
	if got, want := req.Header.Get(headerAPIKey), c.apiKey; got != want {
		t.Errorf("NewRequest() %s is %v, want %v", headerAPIKey, got, want)
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient("", nil)

	type T struct {
		A map[interface{}]interface{}
	}
	_, err := c.NewRequest("GET", ".", &T{})

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error; got %#v.", err)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient("", nil)
	_, err := c.NewRequest("GET", ":", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestNewRequest_errorForNoTrailingSlash(t *testing.T) {
	tests := []struct {
		rawurl    string
		wantError bool
	}{
		{rawurl: "https://example.com/api/v3", wantError: true},
		{rawurl: "https://example.com/api/v3/", wantError: false},
	}
	c := NewClient("", nil)
	for _, test := range tests {
		u, err := url.Parse(test.rawurl)
		if err != nil {
			t.Fatalf("url.Parse returned unexpected error: %v.", err)
		}
		c.BaseURL = u
		if _, err := c.NewRequest(http.MethodGet, "test", nil); test.wantError && err == nil {
			t.Fatalf("Expected error to be returned.")
		} else if !test.wantError && err != nil {
			t.Fatalf("NewRequest returned unexpected error: %v.", err)
		}
	}
}

func TestDo(t *testing.T) {
	type foo struct {
		A string
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"A":"a"}`)
	}))
	defer ts.Close()
	httpClient := ts.Client()
	c := NewClient("", httpClient)
	body := new(foo)
	req, _ := http.NewRequest("GET", ts.URL, nil)
	c.Do(req, body)
	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_invalidJSON(t *testing.T) {
	type foo struct {
		A string
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"A":"a"`)
	}))
	defer ts.Close()
	httpClient := ts.Client()
	c := NewClient("", httpClient)
	body := new(foo)
	req, _ := http.NewRequest("GET", ts.URL, nil)
	_, err := c.Do(req, body)
	if err == nil {
		t.Fatalf("Error expected for malformed JSON but none returned")
	}
}

func TestDo_emptyResponseBody(t *testing.T) {
	type foo struct {
		A string
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()
	httpClient := ts.Client()
	c := NewClient("", httpClient)
	body := new(foo)
	req, _ := http.NewRequest("GET", ts.URL, nil)
	_, err := c.Do(req, body)
	if err != nil {
		t.Fatalf("Empty response body should not trigger an error, instead got %s", err)
	}
}

func TestDo_httpClientError(t *testing.T) {
	httpClient := new(MockHTTPClient)
	httpClient.On(
		"Do",
		&http.Request{},
	).Return(&http.Response{}, errors.New("do_error"))
	c := NewClient("", httpClient)
	_, err := c.Do(&http.Request{}, nil)
	if err == nil {
		t.Fatalf("Error %s not handled", errors.New("do_error"))
	}
}
