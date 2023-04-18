package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST478(email string) (int, error) {
	urlString, _ := url.Parse("https://futureoflondon.us6.list-manage.com/subscribe/post-json?u=9affcec39a&id=c07db14b36&f_id=00c8c2e1f0&c=jQuery19006191744229386273_1680102523853&EMAIL=123dsadsa@gmail.com&FNAME=fasdfasd&LNAME=fdsksdopf&ORG=sdfgsdgsdf&MMERGE5=gsdfgsdfgsd&b_9affcec39a_c07db14b36=&subscribe=Subscribe&_=1680102523854")

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
