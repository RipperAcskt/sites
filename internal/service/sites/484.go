package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST484(email string) (int, error) {
	urlString, _ := url.Parse("https://tennesseetheatre.us12.list-manage.com/subscribe/post-json?u=71d713c05036152b106de6e8a&id=ce75e73885&c=jQuery190023851153975200634_1680176934403&EMAIL=123dsadsa@gmail.com&FNAME=fdasfasd&LNAME=fasdfas&MMERGE3=12332&b_71d713c05036152b106de6e8a_ce75e73885=&subscribe=Subscribe&_=1680176934405")

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
