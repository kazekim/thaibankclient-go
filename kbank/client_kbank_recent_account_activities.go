/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package kbank

import (
	"github.com/mitchellh/mapstructure"
)

func (c *defaultClient) RecentAccountActivities(request *RecentAccountActivitiesRequest) (*RecentAccountActivitiesResponse, error) {

	if c.kBankSvc == nil {
		return nil, ErrKBankConfigNotDefined
	}

	req := c.kBankSvc.NewRecentAccountActivitiesRequest(request)
	resp, err := c.kBankSvc.RecentAccountActivities(req)
	if err != nil {
		return nil, err
	}

	var response RecentAccountActivitiesResponse
	if resp.StatusCode != nil {
		thbErr := NewKBankError(resp.StatusCode, resp.MessageTH, resp.MessageEN)
		response.StatusCode = *resp.StatusCode
		response.Error = thbErr
		return &response, nil
	}

	var activities []RecentAccountActivity

	for _, item := range *resp.Items {

		var activity RecentAccountActivity
		err := mapstructure.Decode(item, &activity)
		if err != nil {
			continue
		}

		activities = append(activities, activity)
	}

	response.StatusCode = "200"
	response.Activities = activities

	return &response, nil
}