/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"fmt"
	"github.com/kazekim/thaibankclient-go/thcerror"
)

type testSSLError struct {
	StatusCode string
	ErrorMsg *string
}

func NewTestSSLError(code, errorMsg *string) thcerror.Error {
	return &testSSLError{
		*code,
		errorMsg,
	}
}

func (e *testSSLError) StatusCodeValue() string {
	return e.StatusCode
}

func (e *testSSLError) MessageTHValue() string {
	return e.MessageTHValue()
}

func (e *testSSLError) MessageENValue() string {
	if e.ErrorMsg != nil {
		return *e.ErrorMsg
	}
	return "-"
}

func (e *testSSLError) Error() string {
	return fmt.Sprintf("code: %v - %v", e.StatusCode, e.MessageENValue())
}