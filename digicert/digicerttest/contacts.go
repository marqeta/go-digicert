package digicerttest

import (
	"github.com/marqeta/go-digicert/digicert"
)

type Contact struct {
	UserID             int    `json:"user_id,omitempty"`
	ContactType        string `json:"contact_type,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	Email              string `json:"email,omitempty"`
	JobTitle           string `json:"job_title,omitempty"`
	Telephone          string `json:"telephone,omitempty"`
	TelephoneExtension string `json:"telephone_extension,omitempty"`
}

func NewOrganizationContact() *digicert.Contact {
	return newContact("organization_contact", "Testorg")
}

func NewTechnicalContact() *digicert.Contact {
	return newContact("technical_contact", "Testtech")
}

func newContact(contact_type, first_name string) *digicert.Contact {
	return &digicert.Contact{
		ContactType: contact_type,
		FirstName:   first_name,
		LastName:    "Contact",
		Email:       "security@marqeta.com",
		JobTitle:    "Security",
		Telephone:   "555-555-5555",
	}
}
