/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"github.com/mitchellh/mapstructure"
)

func (c *defaultClient) KBankCheckBalance(accountNumber string) (*AccountBalance, error) {

	if c.kBankSvc == nil {
		return nil, ErrKBankConfigNotDefined
	}

	req := c.kBankSvc.NewCheckBalanceRequest(accountNumber)
	resp, err := c.kBankSvc.CheckBalance(req)
	if err != nil {
		return nil, err
	}

	var response AccountBalance
	err = mapstructure.Decode(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response,nil
}

