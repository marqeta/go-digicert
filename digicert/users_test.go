package digicert

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func user_test_mock_setup() (*UsersService, *MockClient) {
	s := &UsersService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestUsersGetMe(t *testing.T) {
	s, client := user_test_mock_setup()
	req, _ := http.NewRequest("GET", "user/me", nil)
	client.On(
		"NewRequest",
		"GET",
		"user/me",
		nil,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		new(User),
	).Return(&Response{}, nil).Once()

	s.GetMe()
}

func TestUsersGetMe_newRequestError(t *testing.T) {
	s, client := user_test_mock_setup()
	req, _ := http.NewRequest("GET", "user/me", nil)
	nr_error := errors.New("new_request")
	client.On(
		"NewRequest",
		"GET",
		"user/me",
		nil,
	).Return(req, nr_error).Once()

	_, _, err := s.GetMe()
	if err == nil || !strings.Contains(err.Error(), nr_error.Error()) {
		t.Errorf("Expected error %s, but got %s", nr_error.Error(), err)
	}
}

func TestUsersGetMe_doError(t *testing.T) {
	s, client := user_test_mock_setup()
	req, _ := http.NewRequest("GET", "user/me", nil)
	do_error := errors.New("do")
	client.On(
		"NewRequest",
		"GET",
		"user/me",
		nil,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		new(User),
	).Return(&Response{}, do_error).Once()

	_, _, err := s.GetMe()
	if err == nil || !strings.Contains(err.Error(), do_error.Error()) {
		t.Errorf("Expected error %s, but got %s", do_error.Error(), err)
	}
}

func TestUsersList(t *testing.T) {
	s, client := user_test_mock_setup()
	req, _ := http.NewRequest("GET", "user", nil)
	client.On(
		"NewRequest",
		"GET",
		"user",
		nil,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		new(userList),
	).Return(&Response{}, nil).Once()

	s.List()
}

func TestUsersList_newRequestError(t *testing.T) {
	s, client := user_test_mock_setup()
	req, _ := http.NewRequest("GET", "user", nil)
	nr_error := errors.New("new_request")
	client.On(
		"NewRequest",
		"GET",
		"user",
		nil,
	).Return(req, nr_error).Once()

	_, _, err := s.List()
	if err == nil || !strings.Contains(err.Error(), nr_error.Error()) {
		t.Errorf("Expected error %s, but got %s", nr_error.Error(), err)
	}
}

func TestUsersList_doError(t *testing.T) {
	s, client := user_test_mock_setup()
	req, _ := http.NewRequest("GET", "user", nil)
	do_error := errors.New("do")
	client.On(
		"NewRequest",
		"GET",
		"user",
		nil,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		new(userList),
	).Return(&Response{}, do_error).Once()

	_, _, err := s.List()
	if err == nil || !strings.Contains(err.Error(), do_error.Error()) {
		t.Errorf("Expected error %s, but got %s", do_error.Error(), err)
	}
}

func TestUsersEdit(t *testing.T) {
	var user_id int64 = 1
	s, client := user_test_mock_setup()
	user := &User{ID: user_id}
	req, _ := http.NewRequest("PUT", "user/1", nil)
	client.On(
		"NewRequest",
		"PUT",
		"user/1",
		user,
	).Return(req, nil).Once()
	client.On(
		"Do",
		req,
		user,
	).Return(&Response{}, nil).Once()

	s.Edit(user)
}
