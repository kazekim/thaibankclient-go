/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"github.com/mitchellh/mapstructure"
)

func (c *defaultClient) TestSSL() (*TestSSLResponse, error) {

	if c.kBankSvc == nil {
		return nil, ErrKBankConfigNotDefined
	}

	req, err := c.kBankSvc.NewTestSSLRequest()
	if err != nil {
		return nil, err
	}
	sResp, err := c.kBankSvc.TestSSL(req)
	if err != nil {
		return nil, err
	}

	var response TestSSLResponse
	response.CertificateStatus = sResp.CertificateStatus
	response.PartnerStatus = sResp.PartnerStatus

	if sResp.ErrorMessage != nil {
		thbErr := NewTestSSLError(sResp.StatusCode, sResp.ErrorMessage)
		response.StatusCode = *sResp.StatusCode
		response.Error = thbErr
		return &response, nil
	}

	var certObjects []CertificateObject
	for _, obj := range sResp.CertificateObjects {

		var certObject CertificateObject
		err := mapstructure.Decode(obj, &certObject)
		if err != nil {
			continue
		}

		certObjects = append(certObjects, certObject)
	}

	response.CertificateObjects = certObjects

	return &response,nil
}


