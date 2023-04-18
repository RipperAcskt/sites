package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST480(email string) (int, error) {
	urlString, _ := url.Parse("https://theask.substack.com/subscribe?next=https://theask.substack.com/&later=true&just_signed_up=true&subscription_id=272089092&requires_confirmation=&utm_source=cover_page&email=123asda@gmail.com&skip_redirect_check=true")

	values := urlString.Query()
	values.Set("email", email)
	urlString.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlString.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do failed %s", err.Error())
	}

	return response.StatusCode, nil
}
