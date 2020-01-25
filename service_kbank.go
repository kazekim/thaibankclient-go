/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "github.com/kazekim/req"

const (
	APIKBankCheckBalance = "https://APIPORTAL.kasikornbank.com:12002/deposit/sight/balance/%v"
	)

type kBankService interface {
	CheckBalance(request *serviceKBankCheckBalanceRequest) (*serviceKBankCheckBalanceResponse, error)

	NewCheckBalanceRequest(accountNumber string) *serviceKBankCheckBalanceRequest
}

type defaultKBankService struct {
	req *req.Req
	config *KBankConfig
}

func newKBankService(r *req.Req, config *KBankConfig) kBankService {
	return &defaultKBankService{
		r,
		config,
	}
}