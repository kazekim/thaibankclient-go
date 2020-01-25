/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"fmt"
	"github.com/kazekim/req"
)

type serviceKBankRecentAccountActivitiesRequest struct {
	Header serviceKBankRecentAccountActivitiesRequestHeader
	Param serviceKBankRecentAccountActivitiesRequestParam
}

type serviceKBankRecentAccountActivitiesRequestHeader struct {
	PartnerID string `json:"Partner-Id"`
	PartnerSecret string `json:"Partner-Secret"`
}

type serviceKBankRecentAccountActivitiesRequestParam struct {
	AccountNumber string
}

type serviceKBankRecentAccountActivitiesResponse struct {
	TotalItems int `json:"totalItems"`
	Items []KBankRecentAccountActivity `json:"items`
}

type KBankRecentAccountActivity struct {
	ToAccountNo *string `json:"toAccountNo"`
	ToBankCode *string `json:"toBankCode"`
	ToAccountName *string `json:"toAccountName"`
	ToAccountNameEN *string `json:"toAccountNameEN"`
	FromAccountId *string `json:"fromAccountId"`
	FromBankCode *string `json:"fromBankCode"`
	FromAccountNameTH *string `json:"fromAccountNameTH"`
	FromAccountNameEN *string `json:"fromAccountNameEN"`
	MerchantCode *string `json:"merchantCode"`
	ChannelCode *string `json:"channelCode"`
	ChannelDetail *string `json:"channelDetail"`
	ServiceBranchNo string `json:"serviceBranchNo"`
	ChequeNo string `json:"chequeNo"`
	OutstandingBalance float64 `json:"outstandingBalance"`
	FeeAmount *float64 `json:"feeAmount"`
	TransactionAmount float64 `json:"txnAmount"`
	TransactionDate string `json:"txnDate"`
	TransactionTime string `json:"txnTime"`
	TransactionDesc string `json:"txnDesc"`
	TransactionDescEN string `json:"txnDescEN"`
	EffectiveDate string `json:"effectiveDate"`
	ProxyId *string `json:"proxyId"`
	ProxyTypeCode *string `json:"proxyTypeCode"`
	TellerId string `json:"tellerId"`
	DebitCreditFlag string `json:"debitCreditFlag"`
}

func (s *defaultKBankService) NewRecentAccountActivitiesRequest(accountNumber string) *serviceKBankRecentAccountActivitiesRequest {
	return &serviceKBankRecentAccountActivitiesRequest{
		Header:serviceKBankRecentAccountActivitiesRequestHeader{
			s.config.PartnerID,
			s.config.PartnerSecret,
		},
		Param:serviceKBankRecentAccountActivitiesRequestParam{
			accountNumber,
		},
	}
}
func (s *defaultKBankService) RecentAccountActivities(request *serviceKBankRecentAccountActivitiesRequest) (*serviceKBankRecentAccountActivitiesResponse, error) {

	url := fmt.Sprintf(APIKBankRecentAccountActivities, request.Param.AccountNumber)
	header := req.HeaderFromStruct(request.Header)
	r, err := s.req.Get(url, header, req.Header{
		"cache-control": "no-cache",
	})
	if err != nil {
		return nil, err
	}

	var response serviceKBankRecentAccountActivitiesResponse
	err = r.ToJSON(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}