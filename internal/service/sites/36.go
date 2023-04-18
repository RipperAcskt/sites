package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST36(email string) error {
	urlString := "https://www.atlasobscura.com/aon/email_lists_subscriptions"

	data := url.Values{}
	data.Set("email", email)
	data.Set("authenticity_token", "ClBaRg9KthTX8CwQFpw3ol_KQAFCDcyB8zyjG8UzFI0dWRhz_hI0dCwEWWd4Y33rVZ6lhCjYhjRMcMCshM65sA")
	data.Set("source", "newsletter_landing")
	data.Set("subscribe_daily", "true")
	data.Set("subscribe_weekly", "true")
	data.Set("subscribe_gastro", "true")
	data.Set("subscribe_courses", "true")
	data.Set("subscribe_trips", "true")
	data.Set("commit", "Subscribe")
	encoded := data.Encode()
	req, err := http.NewRequest("POST", urlString, strings.NewReader(encoded))
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do failed %s", err.Error())
	}

	fmt.Println(response.StatusCode)
	return nil
}
