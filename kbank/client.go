/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"errors"
	"github.com/kazekim/req"
	"net/http"
)


var (
	ErrKBankConfigNotDefined = errors.New("kbank config not defined")
)

type Client interface {

	SetHttpClient(client *http.Client)

	CheckBalance(request *CheckBalanceRequest) (*AccountBalanceResponse, error)
	RecentAccountActivities(request *RecentAccountActivitiesRequest) (*RecentAccountActivitiesResponse, error)
	TestSSL() (*TestSSLResponse, error)
}

type defaultClient struct {
	req      *req.Req
	config   Config
	kBankSvc service
}

func NewClient(config Config) Client {

	r := req.New()

	client := defaultClient{
		req: r,
		config: config,
	}

	kBankSvc := newService(r, &config)
	client.kBankSvc = kBankSvc

	return &client
}

func (c *defaultClient) SetHttpClient(client *http.Client) {
	c.req.SetClient(client)
}
