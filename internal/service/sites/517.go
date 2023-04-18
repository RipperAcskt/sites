package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST517(email string) (int, error) {
	urlString, _ := url.Parse("https://assets.mailerlite.com/jsonp/32798/pages/53400036703209426/subscribe?callback=jQuery34107519268847944403_1680197639799&fields[email]=12fdsfd3@gmail.com&ml-submit=1&ajax=1&guid=&_=1680197639800")

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
