package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST899(email string) (int, error) {
	urlString := "https://www.csq.cz/infocentrum/newsletter?tx_mailchimp_registration%5Baction%5D=response&tx_mailchimp_registration%5Bcontroller%5D=Form&cHash=7dbbb5e3a5c63ca675be4c178b9d6a20"

	data := url.Values{}
	data.Set("tx_mailchimp_registration[form][email]", email)
	data.Set("tx_mailchimp_registration[__referrer][@extension]", "Mailchimp")
	data.Set("tx_mailchimp_registration[__referrer][@vendor]", "Sup7even")
	data.Set("Sup7even", "Form")
	data.Set("tx_mailchimp_registration[__referrer][@action]", "index")
	data.Set("tx_mailchimp_registration[__referrer][arguments]", "YTowOnt9f9c4299f02c7217b9f37683b74152a335ac9ed39")

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
