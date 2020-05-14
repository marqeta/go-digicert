package digicert

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func product_test_mock_setup() (*ProductsService, *MockClient) {
	s := &ProductsService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestProductsList(t *testing.T) {
	cases := []struct {
		nrError       error
		DoError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("Do"), errors.New("Do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("GET", "product", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing product listing with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := product_test_mock_setup()
			client.On(
				"NewRequest",
				"GET",
				"product",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(productList),
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.List()
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}
