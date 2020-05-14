package digicert

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func organization_test_mock_setup() (*OrganizationsService, *MockClient) {
	s := &OrganizationsService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestOrganizationsList(t *testing.T) {
	cases := []struct {
		nrError       error
		doError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("do"), errors.New("do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("GET", "organization", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing organization listing with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := organization_test_mock_setup()
			client.On(
				"NewRequest",
				"GET",
				"organization",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(organizationList),
			).Return(&Response{}, c.expectedError).Once()

			_, _, err := s.List()
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestOrganizationsGet(t *testing.T) {
	cases := []struct {
		nrError       error
		doError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("do"), errors.New("do")},
		{nil, nil, nil},
	}
	org := &Organization{ID: 1}
	req, _ := http.NewRequest("GET", "organization/1", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing organization retrieval with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := organization_test_mock_setup()
			client.On(
				"NewRequest",
				"GET",
				"organization/1",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				org,
			).Return(&Response{}, c.expectedError).Once()

			_, _, err := s.Get(1)
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}
