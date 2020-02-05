/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"github.com/kazekim/req"
)

const (
	APIKBankCheckBalance = "/deposit/sight/balance/%v"
	APIKBankRecentAccountActivities = "/deposit/sight/transactions/%v"
	APIKBankTestSSL = "/test/ssl"
	)

type service interface {
	CheckBalance(request *serviceKBankCheckBalanceRequest) (*serviceKBankCheckBalanceResponse, error)
	RecentAccountActivities(request *serviceRecentAccountActivitiesRequest) (*serviceRecentAccountActivitiesResponse, error)
	TestSSL(request *serviceKBankTestSSLRequest) (*serviceKBankTestSSLResponse, error)

	NewCheckBalanceRequest(request *CheckBalanceRequest) *serviceKBankCheckBalanceRequest
	NewRecentAccountActivitiesRequest(request *RecentAccountActivitiesRequest) *serviceRecentAccountActivitiesRequest
	NewTestSSLRequest() (*serviceKBankTestSSLRequest, error)
}

type defaultService struct {
	req *req.Req
	config *Config
}

func newService(r *req.Req, config *Config) service {
	return &defaultService{
		r,
		config,
	}
}