package digicert

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func order_test_mock_setup() (*OrdersService, *MockClient) {
	s := &OrdersService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestOrdersList(t *testing.T) {
	s, client := order_test_mock_setup()
	req, _ := http.NewRequest("GET", "order/certificate", nil)
	client.On(
		"NewRequest",
		"GET",
		"order/certificate",
		nil,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		new(orderList),
	).Return(&Response{}, nil).Once()

	s.List()
}

func TestOrdersList_newRequestError(t *testing.T) {
	s, client := order_test_mock_setup()
	req, _ := http.NewRequest("GET", "order/certificate", nil)
	nr_error := errors.New("new_request")
	client.On(
		"NewRequest",
		"GET",
		"order/certificate",
		nil,
	).Return(req, nr_error).Once()

	_, _, err := s.List()
	if err == nil || !strings.Contains(err.Error(), nr_error.Error()) {
		t.Errorf("Expected error %s, but got %s", nr_error.Error(), err)
	}
}

func TestOrdersList_doError(t *testing.T) {
	s, client := order_test_mock_setup()
	req, _ := http.NewRequest("GET", "order/certificate", nil)
	do_error := errors.New("do")
	client.On(
		"NewRequest",
		"GET",
		"order/certificate",
		nil,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		new(orderList),
	).Return(&Response{}, do_error).Once()

	_, _, err := s.List()
	if err == nil || !strings.Contains(err.Error(), do_error.Error()) {
		t.Errorf("Expected error %s, but got %s", do_error.Error(), err)
	}
}
