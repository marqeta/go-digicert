package digicert

import (
	"fmt"
)

type UsersService service

type User struct {
	ID                      int64        `json:"id,omitempty"`
	Username                string       `json:"username,omitempty"`
	AccountID               int64        `json:"account_id,omitempty"` // CertCentral Account ID
	FirstName               string       `json:"first_name,omitempty"`
	LastName                string       `json:"last_name,omitempty"`
	Email                   string       `json:"email,omitempty"`
	JobTitle                string       `json:"job_title,omitempty"`
	Telephone               string       `json:"telephone,omitempty"`
	Status                  string       `json:"status,omitempty"`
	LastLoginDate           string       `json:"last_login_date,omitempty"`
	IsCertCentral           bool         `json:"is_cert_central,omitempty"`
	IsEnterprise            bool         `json:"is_enterprise,omitempty"`
	IsSAMLSSOOnly           bool         `json:"is_saml_sso_only,omitempty"`
	Type                    string       `json:"type,omitempty"`
	HasContainerAssignments bool         `json:"has_container_assignments,omitempty"`
	AccessRoles             []AccessRole `json:"access_roles,omitempty"`
	Container               *Container   `json:"container,omitempty"`
	ContainerIDAssignments  []int        `json:"container_id_assignments,omitempty"`
}

type AccessRole struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type userList struct {
	Users *[]User
}

func (u User) String() string {
	return Stringify(u)
}

func (s *UsersService) GetMe() (*User, *Response, error) {
	user := new(User)
	resp, err := s.reqHelper("GET", "user/me", nil, user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}

func (s *UsersService) List() (*[]User, *Response, error) {
	list := new(userList)
	resp, err := s.reqHelper("GET", "user", nil, list)
	if err != nil {
		return nil, resp, err
	}
	return list.Users, resp, nil
}

func (s *UsersService) Edit(user *User) (*User, *Response, error) {
	path := fmt.Sprintf("user/%d", user.ID)
	resp, err := s.reqHelper("PUT", path, user, user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}

func (s *UsersService) Create(user *User) (*User, *Response, error) {
	resp, err := s.reqHelper("POST", "user", user, user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}

func (s *UsersService) Delete(user *User) (*User, *Response, error) {
	path := fmt.Sprintf("user/%d", user.ID)
	resp, err := s.reqHelper("DELETE", path, nil, user)
	return nil, resp, err
}

func (s *UsersService) reqHelper(method, path string, body, v interface{}) (*Response, error) {
	req, err := s.client.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, v)
}
