/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "github.com/kazekim/req"

const (
	APIKBankCheckBalance = "/deposit/sight/balance/%v"
	APIKBankRecentAccountActivities = "/deposit/sight/transactions/%v"
	APIKBankTestSSL = "/test/ssl"
	)

type kBankService interface {
	CheckBalance(request *serviceKBankCheckBalanceRequest) (*serviceKBankCheckBalanceResponse, error)
	RecentAccountActivities(request *serviceKBankRecentAccountActivitiesRequest) (*serviceKBankRecentAccountActivitiesResponse, error)
	TestSSL(request *serviceKBankTestSSLRequest) (*serviceKBankTestSSLResponse, error)

	NewCheckBalanceRequest(accountNumber string) *serviceKBankCheckBalanceRequest
	NewRecentAccountActivitiesRequest(accountNumber string) *serviceKBankRecentAccountActivitiesRequest
	NewTestSSLRequest() (*serviceKBankTestSSLRequest, error)
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