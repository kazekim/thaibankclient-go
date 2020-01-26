/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

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

	KBankCheckBalance(accountNumber string) (*AccountBalanceResponse, error)
	KBankRecentAccountActivities(accountNumber string) (*RecentAccountActivitiesResponse, error)
	KBankTestSSL() (*KBankTestSSLResponse, error)
}

type defaultClient struct {
	req *req.Req
	config Config
	kBankSvc kBankService
}

func NewClient(config Config) Client {

	r := req.New()

	client := defaultClient{
		req: r,
		config: config,
	}

	if config.KBank != nil {
		kBankSvc := newKBankService(r, config.KBank)
		client.kBankSvc = kBankSvc
	}

	return &client
}

func (c *defaultClient) SetHttpClient(client *http.Client) {
	c.req.SetClient(client)
}
