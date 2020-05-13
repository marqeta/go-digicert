package digicert

type ContainersService service

type Container struct {
	ID         int    `json:"id,omitempty"`
	PublicID   string `json:"public_id,omitempty"`
	Name       string `json:"name,omitempty"`
	ParentID   int    `json:"parent_id,omitempty"`
	TemplateID int    `json:"template_id,omitempty"`
}
