/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

type KBankRecentAccountActivitiesRequest struct {
	AccountNumber string
}

func NewKBankRecentAccountActivitiesRequest(accountNumber string) *KBankRecentAccountActivitiesRequest {
	return &KBankRecentAccountActivitiesRequest{
		AccountNumber: accountNumber,
	}
}