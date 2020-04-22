package digicert

import (
	"fmt"
	"time"
)

type OrganizationsService service

type Organization struct {
	ID          int64  `json:"id"`
	Status      string `json:"status"`
	Name        string `json:"name"`
	AssumedName string `json:"assumed_name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
	Address     string `json:"address"`
	Address2    string `json:"address2,omitempty"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Telephone   string `json:"telephone,omitempty"`
	Container   struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		IsActive bool   `json:"is_active"`
	} `json:"container"`
	Validations []struct {
		Type           string    `json:"type"`
		Name           string    `json:"name"`
		Description    string    `json:"description"`
		DateCreated    time.Time `json:"date_created,omitempty"`
		ValidatedUntil time.Time `json:"validated_until,omitempty"`
		Status         string    `json:"status"`
		VerifiedUsers  []struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"verified_users,omitempty"`
	} `json:"validations,omitempty"`
	EvApprovers []struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"ev_approvers,omitempty"`
}

type organizationList struct {
	Organizations *[]Organization
}

func (s *OrganizationsService) List() (*[]Organization, *Response, error) {
	list := new(organizationList)
	resp, err := executeAction(s.client, "GET", "organization", nil, list)
	if err != nil {
		return nil, resp, err
	}

	return list.Organizations, resp, nil
}

func (s *OrganizationsService) Get(org_id int64) (*Organization, *Response, error) {
	organization := &Organization{ID: org_id}
	resp, err := executeAction(s.client, "GET", fmt.Sprintf("organization/%d", org_id), nil, organization)
	if err != nil {
		return nil, resp, err
	}
	return organization, resp, nil
}
