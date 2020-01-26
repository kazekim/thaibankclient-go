/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "fmt"

type kBankError struct {
	StatusCode string
	MessageTH *string
	MessageEN *string
}

func NewKBankError(statusCode, messageTH, messageEN *string) THBError {
	return &kBankError{
		*statusCode,
		messageTH,
		messageEN,
	}
}

func (e *kBankError) StatusCodeValue() string {
	return e.StatusCode
}

func (e *kBankError) MessageTHValue() string {
	if e.MessageTH != nil {
		return *e.MessageTH
	}
	return "-"
}

func (e *kBankError) MessageENValue() string {
	if e.MessageEN != nil {
		return *e.MessageEN
	}
	return "-"
}

func (e *kBankError) Error() string {
	return fmt.Sprintf("code: %v - %v (%v)", e.StatusCode, e.MessageENValue(), e.MessageTHValue())
}