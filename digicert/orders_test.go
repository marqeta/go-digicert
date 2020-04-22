package digicert

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func order_test_mock_setup() (*OrdersService, *MockClient) {
	s := &OrdersService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestOrdersList(t *testing.T) {
	cases := []struct {
		nrError       error
		doError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("do"), errors.New("do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("GET", "order/certificate", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing order listing with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := order_test_mock_setup()
			client.On(
				"NewRequest",
				"GET",
				"order/certificate",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(orderList),
			).Return(&Response{}, c.expectedError).Once()

			_, _, err := s.List()
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestOrdersGet(t *testing.T) {
	cases := []struct {
		nrError       error
		doError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("do"), errors.New("do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("GET", "order/certificate/1", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing order retrieval with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := order_test_mock_setup()
			client.On(
				"NewRequest",
				"GET",
				"order/certificate/1",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(Order),
			).Return(&Response{}, c.doError).Once()

			_, _, err := s.Get(1)
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestOrdersCreate(t *testing.T) {
	cases := []struct {
		nrError       error
		doError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("do"), errors.New("do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("POST", "order/certificate/ssl_plus", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing order retrieval with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := order_test_mock_setup()
			newOrder := InitializeOrder()
			client.On(
				"NewRequest",
				"POST",
				"order/certificate/ssl_plus",
				newOrder,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(Order),
			).Return(&Response{}, c.doError).Once()

			_, _, err := s.Create(newOrder)
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestOrdersInitializationDefaults(t *testing.T) {
	defaultOrder := InitializeOrder()

	if defaultOrder.ValidityYears != 1 {
		t.Errorf("Expected default validity period of 1, but got %d", defaultOrder.ValidityYears)
	}

	if defaultOrder.PaymentMethod != "balance" {
		t.Errorf("Expected default payment method of 'balance', but got %s", defaultOrder.PaymentMethod)
	}

	if defaultOrder.SkipApproval {
		t.Error("Expected skip approval to be false, but got true")
	}
}
