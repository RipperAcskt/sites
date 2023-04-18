package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST917(email string) (int, error) {
	urlString := "https://mc.us11.list-manage.com/subscribe/form-post-json?u=7e99bf3837fd69ab443e58213&id=1d177bdfa9&popup=true&EMAIL=qwer%40gmail.com&group%5B30800%5D%5B4096%5D=4096&b_7e99bf3837fd69ab443e58213_1d177bdfa9=&gdpr%5B1037%5D=on&c=dojo_request_script_callbacks.dojo_request_script1"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("popup", "true")
	data.Set("u", "7e99bf3837fd69ab443e58213")
	data.Set("id", "1d177bdfa9")
	data.Set("group[30800][4096]", "4096")
	data.Set("gdpr[1037]", "on")
	data.Set("c", "dojo_request_script_callbacks.dojo_request_script1")
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
