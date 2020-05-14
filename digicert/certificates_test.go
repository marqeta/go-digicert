package digicert

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func certificate_test_mock_setup() (*CertificatesService, *MockClient) {
	s := &CertificatesService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestGetPEMBundle(t *testing.T) {
	cases := []struct {
		nrError       error
		DoError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("Do"), errors.New("Do")},
		{nil, nil, nil},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing PEM retrieval with expected error %s", c.expectedError), func(t *testing.T) {
			bundle := &PEMBundle{PEMs: make([]*PEM, 0)}
			s, client := certificate_test_mock_setup()
			req, _ := http.NewRequest("GET", "certificate/1/chain", nil)
			client.On(
				"NewRequest",
				"GET",
				"certificate/1/chain",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				bundle,
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.GetPEMBundle(1)
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}
