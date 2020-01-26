/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "github.com/mitchellh/mapstructure"

func (c *defaultClient) KBankTestSSL() (*KBankTestSSLResponse, error) {

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

	var response KBankTestSSLResponse
	response.CertificateStatus = sResp.CertificateStatus
	response.PartnerStatus = sResp.PartnerStatus

	if sResp.ErrorMessage != nil {
		thbErr := NewKBankTestSSLError(sResp.StatusCode, sResp.ErrorMessage)
		response.StatusCode = *sResp.StatusCode
		response.Error = thbErr
		return &response, nil
	}

	var certObjects []KBankCertificateObject
	for _, obj := range sResp.CertificateObjects {

		var certObject KBankCertificateObject
		err := mapstructure.Decode(obj, &certObject)
		if err != nil {
			continue
		}

		certObjects = append(certObjects, certObject)
	}

	response.CertificateObjects = certObjects

	return &response,nil
}


