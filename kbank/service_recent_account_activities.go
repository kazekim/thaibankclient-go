/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"fmt"
	"github.com/kazekim/req"
)

type serviceRecentAccountActivitiesRequest struct {
	Url    string
	Header serviceRecentAccountActivitiesRequestHeader
	Param  serviceRecentAccountActivitiesRequestParam
}

type serviceRecentAccountActivitiesRequestHeader struct {
	PartnerID string `json:"Partner-Id"`
	PartnerSecret string `json:"Partner-Secret"`
}

type serviceRecentAccountActivitiesRequestParam struct {
	AccountNumber string
}

type serviceRecentAccountActivitiesResponse struct {
	StatusCode *string             `json:"statusCode"`
	MessageTH *string              `json:"messageTH"`
	MessageEN *string              `json:"messageEN"`
	TotalItems *int                `json:"totalItems"`
	Items *[]RecentAccountActivity `json:"items`
}

func (s *defaultService) NewRecentAccountActivitiesRequest(request *RecentAccountActivitiesRequest) *serviceRecentAccountActivitiesRequest {
	return &serviceRecentAccountActivitiesRequest{
		Url: s.config.BaseUrl+ APIKBankRecentAccountActivities,
		Header: serviceRecentAccountActivitiesRequestHeader{
			s.config.PartnerID,
			s.config.PartnerSecret,
		},
		Param: serviceRecentAccountActivitiesRequestParam{
			request.AccountNumber,
		},
	}
}

func (s *defaultService) RecentAccountActivities(request *serviceRecentAccountActivitiesRequest) (*serviceRecentAccountActivitiesResponse, error) {

	url := fmt.Sprintf(request.Url, request.Param.AccountNumber)
	header := req.HeaderFromStruct(request.Header)
	r, err := s.req.Get(url, header, req.Header{
		"cache-control": "no-cache",
		"content-type": "application/json",
	})
	if err != nil {
		return nil, err
	}

	var response serviceRecentAccountActivitiesResponse
	err = r.ToJSON(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}