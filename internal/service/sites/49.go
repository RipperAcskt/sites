package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST49(email string) error {
	urlString, _ := url.Parse("https://ssl.kpmg.pl/students/api/v1/rejestracja.php?action=sendCode&email=123%40gmail.com&callback=onyx_kpmgstud_rej.jsonp.callback")

	values := urlString.Query()
	values.Set("email", email)
	values.Set("action", "sendCode")
	values.Set("callback", "onyx_kpmgstud_rej.jsonp.callback")
	urlString.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlString.String(), nil)
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

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
