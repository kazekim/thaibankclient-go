/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "fmt"

type kBankError struct {
	Code string
	MessageTH *string
	MessageEN *string
}

func NewKBankError(code, messageTH, messageEN *string) THBError {
	return &kBankError{
		*code,
		messageTH,
		messageEN,
	}
}

func (e *kBankError) CodeValue() string {
	return e.Code
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
	return fmt.Sprintf("code: %v - %v (%v)", e.Code, e.MessageENValue(), e.MessageTHValue())
}