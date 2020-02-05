/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

type RecentAccountActivitiesRequest struct {
	AccountNumber string
}

func NewRecentAccountActivitiesRequest(accountNumber string) *RecentAccountActivitiesRequest {
	return &RecentAccountActivitiesRequest{
		AccountNumber: accountNumber,
	}
}