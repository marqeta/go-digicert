package digicert

import (
	"fmt"
)

type Certificate struct {
	ID                int      `json:"id,omitempty"`
	CommonName        string   `json:"common_name"`
	DNSNames          []string `json:"dns_names,omitempty"`
	CSR               string   `json:"csr"`
	OrganizationUnits []string `json:"organization_units,omitempty"`
	ServerPlatform    struct {
		ID int `json:"id"`
	} `json:"server_platform,omitempty"`
	SignatureHash string `json:"signature_hash"`
	ProfileOption string `json:"profile_option,omitempty"`
	Thumbprint    string `json:"thumbprint,omitempty"`
	SerialNumber  string `json:"serial_number,omitempty"`
}

type PEM struct {
	SubjectCommonName string `json:"subject_common_name"`
	PEM               string `json:"pem"`
}

type PEMBundle struct {
	PEMs []*PEM `json:"intermediates"`
}

func (o Certificate) String() string {
	return Stringify(o)
}

type CertificatesService service

func (s *CertificatesService) GetPEMBundle(cert_id int) (*PEMBundle, *Response, error) {
	bundle := &PEMBundle{PEMs: make([]*PEM, 0)}
	resp, err := executeAction(s.client, "GET", fmt.Sprintf("certificate/%d/chain", cert_id), nil, bundle)
	if err != nil {
		return nil, resp, err
	}
	return bundle, resp, nil
}
