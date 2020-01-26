/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"github.com/mitchellh/mapstructure"
)

func (c *defaultClient) KBankCheckBalance(request *KBankCheckBalanceRequest) (*AccountBalanceResponse, error) {

	if c.kBankSvc == nil {
		return nil, ErrKBankConfigNotDefined
	}

	req := c.kBankSvc.NewCheckBalanceRequest(request)
	resp, err := c.kBankSvc.CheckBalance(req)
	if err != nil {
		return nil, err
	}

	var response AccountBalanceResponse

	if resp.StatusCode != nil {
		thbErr := NewKBankError(resp.StatusCode, resp.MessageTH, resp.MessageEN)
		response.StatusCode = *resp.StatusCode
		response.Error = thbErr
		return &response, nil
	}

	var ab AccountBalance
	err = mapstructure.Decode(resp, &ab)
	if err != nil {
		return nil, err
	}
	response.StatusCode = "200"
	response.AccountBalance = &ab

	return &response,nil
}

