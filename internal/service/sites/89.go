package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST89(email string) (int, error) {
	urlString := "https://youthenergy.eu/subscribe-newsletter/"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("FNAME", "qwer")
	data.Set("LNAME", "qwer")
	data.Set("MMERGE6", "qwer")
	data.Set("MMERGE3", "qwer")
	data.Set("_mc4wp_form_element_id", "qwer")
	data.Set("_mc4wp_form_id", "3090")
	data.Set("_mc4wp_timestamp", "1680129576")
	data.Set("AGREE_TO_TERMS", "1")

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
