package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST479(email string) (int, error) {
	urlString, _ := url.Parse("https://immunology.us17.list-manage.com/subscribe/post-json?u=ba9edbc1f3f5afa275c114464&id=98bf7bdea8&c=jQuery19008693475069110812_1680102658057&EMAIL=123dsadsa@gmail.com&group[37113]=2&interestgroup_field=&gdpr[85225]=Y&b_ba9edbc1f3f5afa275c114464_98bf7bdea8=&subscribe=Subscribe&_=1680102658058")

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
