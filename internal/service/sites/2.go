package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST2(email string) error {
	urlString := "https://a.klaviyo.com/client/subscriptions/?company_id=YdHW8e"

	data := fmt.Sprintf(`{
		"data": {
			"type": "subscription",
			"attributes": {
				"list_id": "SULFf3",
				"custom_source": "Newsletter Signup",
				"email": "%s",
				"properties": {
					"marketing_consent": [
						"Yes, I would like to receive email marketing communication from Later. I understand that I can unsubscribe at any time."
					],
					"$consent_method": "Klaviyo Form",
					"$consent_form_id": "WUdRKB",
					"$consent_form_version": 4370049,
					"$timezone_offset": 3,
					"$exchange_id": "Hrqg8YUkYSK7r8ewNj4yrmXsj5EH0nKjNn3nHuIn0_k=.YdHW8e"
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
