/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"fmt"
	"github.com/kazekim/req"
)

type serviceKBankCheckBalanceRequest struct {
	Url string
	Header serviceKBankCheckBalanceRequestHeader
	Param serviceKBankCheckBalanceRequestParam
}

type serviceKBankCheckBalanceRequestHeader struct {
	PartnerID string `json:"Partner-Id"`
	PartnerSecret string `json:"Partner-Secret"`
}

type serviceKBankCheckBalanceRequestParam struct {
	AccountNumber string
}

type serviceKBankCheckBalanceResponse struct {
	StatusCode *string `json:"statusCode"`
	MessageTH *string  `json:"messageTH"`
	MessageEN *string  `json:"messageEN"`
	AvailableBalance *float64 `json:"availBalance"`
	AccountBalance *float64 `json:"acctBalance"`
	AccountStatus *string `json:"acctStatus"`
}

func (s *defaultKBankService) NewCheckBalanceRequest(request *KBankCheckBalanceRequest) *serviceKBankCheckBalanceRequest {
	return &serviceKBankCheckBalanceRequest{
		Url: s.config.BaseUrl+APIKBankCheckBalance,
		Header:serviceKBankCheckBalanceRequestHeader{
			s.config.PartnerID,
			s.config.PartnerSecret,
		},
		Param:serviceKBankCheckBalanceRequestParam{
			request.AccountNumber,
		},
	}
}

func (s *defaultKBankService) CheckBalance(request *serviceKBankCheckBalanceRequest) (*serviceKBankCheckBalanceResponse, error) {

	url := fmt.Sprintf(request.Url, request.Param.AccountNumber)
	header := req.HeaderFromStruct(request.Header)
	r, err := s.req.Get(url, header, req.Header{
		"cache-control": "no-cache",
		"content-type": "application/json",
	})
	if err != nil {
		return nil, err
	}

	var response serviceKBankCheckBalanceResponse
	err = r.ToJSON(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
