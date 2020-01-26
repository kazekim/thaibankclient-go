/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

type KBankConfig struct {
	BaseUrl string
	PartnerID string `json:"partner_id"`
	PartnerSecret string `json:"partner_secret"`
}
