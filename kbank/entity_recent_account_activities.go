/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"github.com/kazekim/thaibankclient-go/thcerror"
)

type RecentAccountActivitiesResponse struct {
	StatusCode string
	Activities []RecentAccountActivity
	Error      thcerror.Error
}

type RecentAccountActivity struct {
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