package digicerttest

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/marqeta/go-digicert/digicert"
)

func NewTestCertificate(common_name string) (*digicert.Certificate, error) {
	dcert := &digicert.Certificate{}
	template := &x509.CertificateRequest{}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	csr, err := x509.CreateCertificateRequest(rand.Reader, template, key)
	if err != nil {
		return nil, err
	}
	csr_text := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE REQUEST",
			Bytes: csr,
		},
	)
	dcert.CSR = fmt.Sprintf("%s", csr_text)
	dcert.CommonName = common_name
	dcert.ServerPlatform.ID = -1
	dcert.DNSNames = []string{common_name}
	dcert.SignatureHash = "sha256"
	dcert.OrganizationUnits = []string{"Certificate Factory Test Org Inc LLC"}

	return dcert, nil
}
