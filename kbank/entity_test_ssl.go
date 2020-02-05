/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"github.com/kazekim/thaibankclient-go/thcerror"
)

type TestSSLResponse struct {
	StatusCode         string
	CertificateStatus  *string
	PartnerStatus      *string
	CertificateObjects []CertificateObject
	Error              thcerror.Error
}

type CertificateObject struct {
	Subject string
	Issuer string
	StartDate string
	ExpireDate string
	SerialNumber string
}
