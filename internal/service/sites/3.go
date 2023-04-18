package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST3(email string) error {
	urlString := "https://a.klaviyo.com/client/subscriptions/?company_id=MTNfzL"

	data := fmt.Sprintf(`{
		"data": {
			"type": "subscription",
			"attributes": {
				"list_id": "PxySrP",
				"custom_source": "Sign up to our newsletter - Newsletter Subscription Page",
				"email": "%s",
				"properties": {
					"_admin_source": "Newsletter subscription",
					"$consent": "email",
					"consent": "email",
					"Accepts Marketing": "True",
					"Privacy Policy": "YES",
					"$consent_method": "Klaviyo Form",
					"$consent_form_id": "JvB2Uy",
					"$consent_form_version": 6712326,
					"services": "{\"shopify\":{\"source\":\"form\"}}",
					"$timezone_offset": 3,
					"$exchange_id": "PywMuOZtOkRpA-xmGGRO-UO5Cgd9dWno2wpQgKhFN1o=.MTNfzL"
				}
			}
		}
	}`, email)

	req, err := http.NewRequest("POST", urlString, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("revision", "2022-02-16")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do failed %s", err.Error())
	}

	fmt.Println(response.StatusCode)
	return nil
}
