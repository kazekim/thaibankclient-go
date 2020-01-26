/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/kazekim/thaibankclient-go"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	g, err := thaibankclient.NewTLSCertGenerator()
	if err != nil {
		panic(err)
	}

	keyPair, err := g.GenerateDefaultX509KeyPair()
	if err != nil {
		panic(err)
	}

	requestBody, err := json.Marshal(map[string]string{
		"partnerId":     "PTR7978721",
		"partnerSecret": "7a053e89d8364030b3501f7956b6e56b",
	})
	if err != nil {
		log.Fatalln(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{*keyPair},
		},
	}
	url := "https://APIPORTAL.kasikornbank.com:12002/test/ssl"
	client := &http.Client{Transport: tr}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("verify result", string(body))
}
