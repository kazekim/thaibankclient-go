/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"github.com/kazekim/thaibankclient-go/thcerror"
)

const (
	Active AccountStatus = "Active"
	Unknown AccountStatus = "Unknown"
)

type AccountStatus string

type AccountBalanceResponse struct {
	StatusCode     string
	AccountBalance *AccountBalance
	Error          thcerror.Error
}

type AccountBalance struct {
	AvailableBalance float64
	AccountBalance float64
	AccountStatus string
	Error *thcerror.Error
}

func (ab *AccountBalance) Status() AccountStatus {

	switch ab.AccountStatus {
	case string(Active):
		return Active
	default:
		return Unknown
	}
}