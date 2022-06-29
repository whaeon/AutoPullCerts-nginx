package app

import "testing"

func TestDownload(t *testing.T) {
	cert := NewRequestParams(false, "", 0, 20, 1)
	domain := "www.example.com"
	CertificateId := "xxxxxx"
	emptyStr := ""
	certificate := NewCertInfo(&CertificateId, &domain, &emptyStr, &emptyStr)
	certs = append(certs, certificate)
	cert.Download()
}
