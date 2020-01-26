/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient


type KBankTestSSLResponse struct {
	StatusCode         string
	CertificateStatus  *string
	PartnerStatus      *string
	CertificateObjects []KBankCertificateObject
	Error              THBError
}

type KBankCertificateObject struct {
	Subject string
	Issuer string
	StartDate string
	ExpireDate string
	SerialNumber string
}
