/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import "github.com/mitchellh/mapstructure"

func (c *defaultClient) KBankRecentAccountActivities(accountNumber string) (*[]RecentAccountActivity, error) {

	if c.kBankSvc == nil {
		return nil, ErrKBankConfigNotDefined
	}

	req := c.kBankSvc.NewRecentAccountActivitiesRequest(accountNumber)
	resp, err := c.kBankSvc.RecentAccountActivities(req)
	if err != nil {
		return nil, err
	}

	var activities []RecentAccountActivity

	for _, item := range resp.Items {

		var activity RecentAccountActivity
		err := mapstructure.Decode(item, &activity)
		if err != nil {
			continue
		}

		activities = append(activities, activity)
	}

	return &activities, nil
}

