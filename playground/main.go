/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/kazekim/thaibankclient-go"
)

func main() {

	config := thaibankclient.Config{
		KBank: &thaibankclient.KBankConfig{
			"PTR7978721",
			"7a053e89d8364030b3501f7956b6e56b",
		},
	}

	c := thaibankclient.NewClient(config)
	accountBalance, err := c.KBankCheckBalance("1246796189")
	if err != nil {
		panic(err)
	}

	Print(accountBalance)

	recentActivities, err := c.KBankRecentAccountActivities("1246796189")
	if err != nil {
		panic(err)
	}

	Print(recentActivities)
}


func Print(value interface{}) {
	marshalStruct, _ := json.MarshalIndent(value, "", "\t")
	fmt.Println(string(marshalStruct))
}