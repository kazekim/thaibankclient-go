/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"crypto/tls"
	"encoding/json"
	"github.com/kazekim/req"
	"net/http"
	"fmt"
)

type serviceKBankTestSSLRequest struct {
	Url string
	Body serviceKBankTestSSLRequestBody
	TLSCertKeyPair *tls.Certificate
}

type serviceKBankTestSSLRequestBody struct {
	PartnerID string `json:"partnerId"`
	PartnerSecret string `json:"partnerSecret"`
}

type serviceKBankTestSSLResponse struct {
	StatusCode *string `json:"statusCode"`
	CertificateStatus *string  `json:"CertStatus"`
	PartnerStatus *string  `json:"partnerStatus"`
	ErrorMessage *string `json:"errorMsg"`
	CertificateObjects []serviceKBankCertificateObjects `json:"certificateObjs`
}

type serviceKBankCertificateObjects struct {
	Subject string `json:"subject"`
	Issuer string `json:"issuer"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	SerialNumber string `json:"serialNumber"`
}

func (s *defaultKBankService) NewTestSSLRequest() (*serviceKBankTestSSLRequest, error) {

	g, err := NewTLSCertGenerator()
	if err != nil {
		return nil, err
	}

	keyPair, err := g.GenerateDefaultX509KeyPair()
	if err != nil {
		return nil, err
	}

	return &serviceKBankTestSSLRequest{
		Url: s.config.BaseUrl+APIKBankTestSSL,
		Body:serviceKBankTestSSLRequestBody{
			s.config.PartnerID,
			s.config.PartnerSecret,
		},
		TLSCertKeyPair: keyPair,
	}, nil
}

func (s *defaultKBankService) TestSSL(request *serviceKBankTestSSLRequest) (*serviceKBankTestSSLResponse, error) {

	url := request.Url
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{*request.TLSCertKeyPair},
		},
	}
	client := &http.Client{Transport: tr}
	r, err := s.req.Post(
		url,
		req.Header{
			"content-type": "application/json",
		},
		req.BodyJSON(&request.Body),
		client,
	)
	if err != nil {
		return nil, err
	}
	//
	//body, err := ioutil.ReadAll(r.Response().Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Println("verify result", string(body))

	var response serviceKBankTestSSLResponse
	err = r.ToJSON(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Print(value interface{}) {
	marshalStruct, _ := json.MarshalIndent(value, "", "\t")
	fmt.Println(string(marshalStruct))
}