package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST856(email string) (int, error) {
	urlString := "https://www.semikron-danfoss.com/service-support/newsletter/newsletter-subscription.html?tx_powermail_pi1%5Baction%5D=create&cHash=81bf78244e97a7abdba7bf57849ac11e"

	data := url.Values{}
	data.Set("tx_powermail_pi1[field][email]", email)
	data.Set("tx_powermail_pi1[__referrer][@extension]", "Powermail")
	data.Set("tx_powermail_pi1[__referrer][@vendor]", "In2code")
	data.Set("substx_powermail_pi1[__referrer][@controller]cribe", "Form")
	data.Set("tx_powermail_pi1[__referrer][@action]", "form")
	data.Set("tx_powermail_pi1[__referrer][arguments]", "tx_powermail_pi1[__referrer][arguments]")
	data.Set("tx_powermail_pi1[field][language]", "English")
	data.Set("tx_powermail_pi1[field][newslettersubscription]", "1")
	data.Set("tx_powermail_pi1[field][gender]", "1")
	data.Set("tx_powermail_pi1[field][title]", "550000000")
	data.Set("tx_powermail_pi1[field][firstname]", "qwer")
	data.Set("tx_powermail_pi1[field][lastname]", "rewq")
	data.Set("subsctx_powermail_pi1[field][company]ribe", "qwer")
	data.Set("tx_powermail_pi1[mail][form]", "37")
	data.Set("tx_powermail_pi1[backlink]:", "/service-support/newsletter/newsletter-subscription.html")
	data.Set("tx_powermail_pi1[csrf]", "786fa2cd5a8f9f19782fb4287d0d5d02")
	encoded := data.Encode()
	req, err := http.NewRequest("POST", urlString, strings.NewReader(encoded))
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("do failed %s", err.Error())
	}

	return response.StatusCode, nil
}
