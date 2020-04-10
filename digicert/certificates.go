package digicert

type Certificate struct {
	ID                int64    `json:"id"`
	CommonName        string   `json:"common_name"`
	DNSNames          []string `json:"dns_names"`
	CSR               string   `json:"csr"`
	OrganizationUnits []string `json:"organization_units,omitempty"`
	ServerPlatform    struct {
		ID int `json:"id"`
	} `json:"server_platform,omitempty"`
	SignatureHash string `json:"signature_hash"`
	ProfileOption string `json:"profile_option,omitempty"`
}

func (o Certificate) String() string {
	return Stringify(o)
}
