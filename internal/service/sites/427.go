package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST427(email string) (int, error) {
	urlString := "https://housing.lacity.org/about-us/subscribe-to-newsletters"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("FNAME", "fspadjfpijas")
	data.Set("LNAME", "jfdoisjaiofasd")
	data.Set("_mc4wp_lists[]", "d20255fa95")
	data.Set("_mc4wp_lists[]", "3f6ceb3592")
	data.Set("_mc4wp_lists[]", "4f2fbdc17d")
	data.Set("_mc4wp_lists[]", "ba4d74dc18")
	data.Set("_mc4wp_lists[]", "7723c1a44a")
	data.Set("_mc4wp_lists[]", "4ae83d5066")
	data.Set("_mc4wp_lists[]", "aea938a7c2")
	data.Set("_mc4wp_lists[]", "f1df68d0b0")
	data.Set("_mc4wp_lists[]", "63f8c7352b")
	data.Set("_mc4wp_lists[]", "364b3a2f55")
	data.Set("_mc4wp_lists[]", "a3b7401602")
	data.Set("_mc4wp_timestamp", "1680015383")
	data.Set("_mc4wp_form_id", "6142")
	data.Set("_mc4wp_timestamp", "1680015383")
	data.Set("_mc4wp_form_element_id", "mc4wp-form-1")

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
