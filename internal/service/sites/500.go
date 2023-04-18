package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST500(email string) (int, error) {
	urlString, _ := url.Parse("https://assets.mailerlite.com/jsonp/48000/forms/54945959515260455/subscribe?callback=jQuery183009616611822222887_1680191613992&fields[email]=123fdasf@gmail.com&ml-submit=1&anticsrf=true&ajax=1&guid=27fc0788-3b75-9e6a-cab8-a5b8d4a7fd37&_=1680191633270")

	values := urlString.Query()
	values.Set("fields[email]", email)
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
