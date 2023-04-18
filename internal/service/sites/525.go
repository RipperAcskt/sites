package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST525(email string) (int, error) {
	urlString, _ := url.Parse("https://acenapdx.us2.list-manage.com/subscribe/post-json?u=e9c25161edc19c537ced8aec6&id=7aa26c7aa5&c=jQuery1900040437062201496454_1680201368376&EMAIL=123dfasdsdfas@gmail.com&FNAME=afasd&LNAME=fsadfasdf&b_e9c25161edc19c537ced8aec6_7aa26c7aa5=&subscribe=Subscribe&_=1680201368377")

	values := urlString.Query()
	values.Set("EMAIL", email)
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
