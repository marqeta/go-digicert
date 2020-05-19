package digicerttest

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
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

func NewChainBundle(chain_length int) (*digicert.PEMBundle, error) {
	if chain_length <= 0 {
		return nil, errors.New("chain_length must be a positive integer value")
	}
	bundle := &digicert.PEMBundle{PEMs: make([]*digicert.PEM, chain_length)}
	for i, _ := range bundle.PEMs {
		bundle.PEMs[i] = newPEM("test")
	}
	return bundle, nil
}

func newPEM(subject_common_name string) *digicert.PEM {
	// NOTE in the future we can make these attributes more useful, e.g. set a real PEM value
	return &digicert.PEM{SubjectCommonName: subject_common_name}
}
