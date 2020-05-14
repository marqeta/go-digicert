package digicert

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
