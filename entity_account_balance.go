/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

const (
	Active AccountStatus = "Active"
	Unknown AccountStatus = "Unknown"
)

type AccountStatus string

type AccountBalance struct {
	AvailableBalance float64
	AccountBalance float64
	AccountStatus string
}

func (ab *AccountBalance) Status() AccountStatus {

	switch ab.AccountStatus {
	case string(Active):
		return Active
	default:
		return Unknown
	}
}