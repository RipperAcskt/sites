package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST95(email string) (int, error) {
	urlString := "https://allthingsiceland.us15.list-manage.com/subscribe/post"

	data := url.Values{}
	data.Set("MERGE0", email)
	data.Set("u", "2ef8d6d11617b848a02545ba2")
	data.Set("id", "360a92fe00")
	data.Set("MERGE1", "qwer")
	data.Set("submit", "Subscribe")
	data.Set("ht", "e62171df8f53090d47bff4dd4e9726a3c530aefb:MTY4MDEzMTEyMS40Njg3")
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
