package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST86(email string) (int, error) {
	urlString := "https://uni-foundation.us6.list-manage.com/subscribe/post"

	data := url.Values{}
	data.Set("MERGE0", email)
	data.Set("u", "21c856853cf445ddff1c67120")
	data.Set("id", "240b7b3384")
	data.Set("MERGE1", "qwer")
	data.Set("MERGE2", "qwer")
	data.Set("gdpr[233]", "Y")
	data.Set("gdpr[865]", "Y")
	data.Set("submit", "Subscribe")
	data.Set("ht", "35d415385d0f97cb30db2ff5659afd47630c8724:MTY4MDEyOTA1Mi4zNjc=")
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
