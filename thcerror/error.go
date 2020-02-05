/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thcerror

type Error interface {
	StatusCodeValue() string
	MessageTHValue() string
	MessageENValue() string
	Error() string
}
