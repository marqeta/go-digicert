package digicert

type ContainersService service

type Container struct {
	ID         int64  `json:"id,omitempty"`
	PublicID   string `json:"public_id,omitempty"`
	Name       string `json:"name,omitempty"`
	ParentID   int64  `json:"parent_id,omitempty"`
	TemplateID int64  `json:"template_id,omitempty"`
}
