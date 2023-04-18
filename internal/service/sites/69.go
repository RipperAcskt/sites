package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST69(email string) (int, error) {
	urlString := "https://ontario.us16.list-manage.com/subscribe/post"

	data := url.Values{}
	data.Set("MERGE0", email)
	data.Set("u", "4ebaa11387038c3e29a6a4a8b")
	data.Set("id", "5f2bc2b632")
	data.Set("MERGE1", "qwer")
	data.Set("MERGE2", "qwer")
	data.Set("submit", "Subscribe")
	data.Set("ht", "6e5632827f67f1d773707f6ccf498f7650248877:MTY4MDEyNTg0OC44NzAx")
	data.Set("mc_signupsource", "hosted")
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
