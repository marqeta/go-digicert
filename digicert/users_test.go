package digicert

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func user_test_mock_setup() (*UsersService, *MockClient) {
	s := &UsersService{}
	client := &MockClient{}
	s.client = client
	return s, client
}

func TestUsersGetMe(t *testing.T) {
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
		t.Run(fmt.Sprintf("Testing self retrieval with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := user_test_mock_setup()
			req, _ := http.NewRequest("GET", "user/me", nil)
			client.On(
				"NewRequest",
				"GET",
				"user/me",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(User),
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.GetMe()
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestUsersGet(t *testing.T) {
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
		t.Run(fmt.Sprintf("Testing self retrieval with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := user_test_mock_setup()
			req, _ := http.NewRequest("GET", "user/1", nil)
			client.On(
				"NewRequest",
				"GET",
				"user/1",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(User),
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.Get(1)
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestUsersList(t *testing.T) {
	cases := []struct {
		nrError       error
		DoError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("Do"), errors.New("Do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("GET", "user", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing user listing with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := user_test_mock_setup()
			client.On(
				"NewRequest",
				"GET",
				"user",
				nil,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				new(userList),
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.List()
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestUsersEdit(t *testing.T) {
	user := &User{ID: 1}
	cases := []struct {
		nrError       error
		DoError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("Do"), errors.New("Do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("GET", "user/1", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing user edits with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := user_test_mock_setup()
			client.On(
				"NewRequest",
				"PUT",
				"user/1",
				user,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				user,
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.Edit(user)
			testExpectedErrorChecker(t, c.expectedError, err)
		})
	}
}

func TestUsersCreate(t *testing.T) {
	user := &User{}
	cases := []struct {
		nrError       error
		DoError       error
		expectedError error
	}{
		{errors.New("new_request"), nil, errors.New("new_request")},
		{nil, errors.New("Do"), errors.New("Do")},
		{nil, nil, nil},
	}
	req, _ := http.NewRequest("POST", "user", nil)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing user creation with expected error %s", c.expectedError), func(t *testing.T) {
			s, client := user_test_mock_setup()
			client.On(
				"NewRequest",
				"POST",
				"user",
				user,
			).Return(req, c.nrError).Once()
			client.On(
				"Do",
				req,
				user,
			).Return(&Response{}, c.DoError).Once()

			_, _, err := s.Create(user)
			testExpectedErrorChecker(t, c.expectedError, err)

		})
	}
}
