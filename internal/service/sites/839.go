package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST839(email string) (int, error) {
	urlString := "https://annfriedman.us7.list-manage.com/subscribe/post-json?u=4a77dae67a768bc3b920d4961&id=e8184bf53a&c=jQuery19006157713928957314_1679939670243&EMAIL=qwer%40gmail.com&MMERGE1=qwer&b_4a77dae67a768bc3b920d4961_e8184bf53a=&subscribe=Subscribe&_=1679939670244"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("MMERGE1", "qwer")
	data.Set("subscribe", "Subscribe")
	data.Set("u", "4a77dae67a768bc3b920d4961")
	data.Set("id", "e8184bf53a")
	data.Set("c", "jQuery19006157713928957314_1679939670243")
	data.Set("_", "1679939670244")
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
