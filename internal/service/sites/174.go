package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST174(email string) (int, error) {
	urlString := "https://www.osc.ca/fr/nouvelles-evenements/abonnez-vous"

	data := url.Values{}
	data.Set("mergevars[EMAIL]", email)
	data.Set("mergevars[FNAME]", "qwer")
	data.Set("mergevars[LNAME]", "qwer")
	data.Set("mergevars[ORG]", "qwer")
	data.Set("form_build_id", "form-ysjuEffX1nU-k7Ezc4Flse76RkBHCF8YWaQo4NAksK4")
	data.Set("form_id", "mailchimp_signup_subscribe_block_osc_email_alerts_signup_form_form")
	data.Set("mailchimp_lists[interest_groups][b46fd26fa4][df4ff4314f]", "df4ff4314f")
	data.Set("op", "Subscribe")
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
