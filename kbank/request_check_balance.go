/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

type CheckBalanceRequest struct {
	AccountNumber string
}

func NewCheckBalanceRequest(accountNumber string) *CheckBalanceRequest {
	return &CheckBalanceRequest{
		AccountNumber: accountNumber,
	}
}