/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

type KBankCheckBalanceRequest struct {
	AccountNumber string
}

func NewKBankCheckBalanceRequest(accountNumber string) *KBankCheckBalanceRequest {
	return &KBankCheckBalanceRequest{
		AccountNumber: accountNumber,
	}
}