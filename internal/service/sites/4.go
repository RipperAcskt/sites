package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST4(email string) error {
	urlString, _ := url.Parse("https://api.cgtn.com/sub/newsletter/api/newsletters/email/subscribe?topicId=3&email=123%40gmail.com")

	values := urlString.Query()
	values.Set("email", email)
	values.Set("topicId", fmt.Sprint(1))
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
