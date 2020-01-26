/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

type RecentAccountActivitiesResponse struct {
	StatusCode string
	Activities []RecentAccountActivity
	Error THBError
}

type RecentAccountActivity struct {
	ToAccountNo *string
	ToBankCode *string
	ToAccountName *string
	ToAccountNameEN *string
	FromAccountId *string
	FromBankCode *string
	FromAccountNameTH *string
	FromAccountNameEN *string
	MerchantCode *string
	ChannelCode *string
	ChannelDetail *string
	ServiceBranchNo string
	ChequeNo string
	OutstandingBalance float64
	FeeAmount *float64
	TransactionAmount float64
	TransactionDate string
	TransactionTime string
	TransactionDesc string
	TransactionDescEN string
	EffectiveDate string
	ProxyId *string
	ProxyTypeCode *string
	TellerId string
	DebitCreditFlag string
}
