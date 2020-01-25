/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"fmt"
	"github.com/kazekim/req"
)

type serviceKBankCheckBalanceRequest struct {
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
	AvailableBalance float64 `json:"availBalance"`
	AccountBalance float64 `json:"acctBalance"`
	AccountStatus string `json:"acctStatus"`
}

func (s *defaultKBankService) NewCheckBalanceRequest(accountNumber string) *serviceKBankCheckBalanceRequest {
	return &serviceKBankCheckBalanceRequest{
		Header:serviceKBankCheckBalanceRequestHeader{
			s.config.PartnerID,
			s.config.PartnerSecret,
		},
		Param:serviceKBankCheckBalanceRequestParam{
			accountNumber,
		},
	}
}

func (s *defaultKBankService) CheckBalance(request *serviceKBankCheckBalanceRequest) (*serviceKBankCheckBalanceResponse, error) {

	url := fmt.Sprintf(APIKBankCheckBalance, request.Param.AccountNumber)
	header := req.HeaderFromStruct(request.Header)
	r, err := s.req.Get(url, header, req.Header{
		"cache-control": "no-cache",
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