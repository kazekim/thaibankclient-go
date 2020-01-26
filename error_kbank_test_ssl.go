/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"fmt"
)

type kBankTestSSLError struct {
	StatusCode string
	ErrorMsg *string
}

func NewKBankTestSSLError(code, errorMsg *string) THBError {
	return &kBankTestSSLError{
		*code,
		errorMsg,
	}
}

func (e *kBankTestSSLError) StatusCodeValue() string {
	return e.StatusCode
}

func (e *kBankTestSSLError) MessageTHValue() string {
	return e.MessageTHValue()
}

func (e *kBankTestSSLError) MessageENValue() string {
	if e.ErrorMsg != nil {
		return *e.ErrorMsg
	}
	return "-"
}

func (e *kBankTestSSLError) Error() string {
	return fmt.Sprintf("code: %v - %v", e.StatusCode, e.MessageENValue())
}