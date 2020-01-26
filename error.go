/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

type THBError interface {
	CodeValue() string
	MessageTHValue() string
	MessageENValue() string
	Error() string
}
