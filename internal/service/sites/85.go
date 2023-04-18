package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST85(email string) (int, error) {
	urlString := "https://european-agency.us6.list-manage.com/subscribe/post?u=f5d5c0af5ae79d814d3545261&id=66e0e53f5b"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("FNAME", "qwer")
	data.Set("LNAME", "qwer")
	data.Set("MMERGE3", "qwer")
	data.Set("EMAILTYPE", "html")
	data.Set("subscribe", "Subscribe")
	data.Set("u", "f5d5c0af5ae79d814d3545261")
	data.Set("id", "66e0e53f5b")
	encoded := data.Encode()
	req, err := http.NewRequest("POST", urlString, strings.NewReader(encoded))
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("do failed %s", err.Error())
	}

	return response.StatusCode, nil
}
