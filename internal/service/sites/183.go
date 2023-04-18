package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST183(email string) (int, error) {
	urlString := "https://www.so-ups.ru/news/subscription/by-email/?tx_msemailsubscription_pi%5Baction%5D=new&tx_msemailsubscription_pi%5Bcontroller%5D=Subscription&cHash=02948c50d1d999bddda119e60d4be180"

	data := url.Values{}
	data.Set("tx_msemailsubscription_pi[subscription][email]", email)
	data.Set("tx_msemailsubscription_pi[subscription][newsCategoriesArray][]", "2")
	data.Set("tx_msemailsubscription_pi[subscription][newsCategoriesArray][]", "3")
	data.Set("tx_msemailsubscription_pi[__referrer][@extension]", "MsEmailSubscription")
	data.Set("tx_msemailsubscription_pi[__referrer][@vendor]", "MS")
	data.Set("tx_msemailsubscription_pi[__referrer][@controller]", "Subscription")
	data.Set("tx_msemailsubscription_pi[__referrer][@action]", "new")
	data.Set("tx_msemailsubscription_pi[__referrer][arguments]", "YTowOnt99cbc1058fccf67d12f591a05a95ea92cb99eb7b2")
	data.Set("tx_msemailsubscription_pi[__referrer][@request]", `a:4:{s:10:"@extension";s:19:"MsEmailSubscription";s:11:"@controller";s:12:"Subscription";s:7:"@action";s:3:"new";s:7:"@vendor";s:2:"MS";}df36cdc17fb5d95bcaf2578e54cb4bb34af8f885`)
	data.Set("tx_msemailsubscription_pi[action]", "new")
	data.Set("tx_msemailsubscription_pi[controller]", "Subscription")
	data.Set("cHash", "02948c50d1d999bddda119e60d4be180")
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
