/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

type THBError interface {
	StatusCodeValue() string
	MessageTHValue() string
	MessageENValue() string
	Error() string
}
