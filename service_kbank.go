/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "github.com/kazekim/req"

const (
	APIKBankCheckBalance = "https://APIPORTAL.kasikornbank.com:12002/deposit/sight/balance/%v"
	APIKBankRecentAccountActivities = "https://APIPORTAL.kasikornbank.com:12002/deposit/sight/transactions/%v"
	)

type kBankService interface {
	CheckBalance(request *serviceKBankCheckBalanceRequest) (*serviceKBankCheckBalanceResponse, error)
	RecentAccountActivities(request *serviceKBankRecentAccountActivitiesRequest) (*serviceKBankRecentAccountActivitiesResponse, error)

	NewCheckBalanceRequest(accountNumber string) *serviceKBankCheckBalanceRequest
	NewRecentAccountActivitiesRequest(accountNumber string) *serviceKBankRecentAccountActivitiesRequest
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