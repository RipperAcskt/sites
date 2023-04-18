package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST952(email string) (int, error) {
	urlString := "https://austrian.audio/an-englishman-in-austria-sting-to-perform-in-eisenstadt/?cc_mailchimp_ajax_func=cc_mailchimp_sub_register"

	data := url.Values{}
	data.Set("cc_email", email)
	data.Set("cc_name", "qwer")
	data.Set("cc_country", "United Kingdom")
	data.Set("cc_tnb", "Y")
	data.Set("cc_page", "https://austrian.audio/an-englishman-in-austria-sting-to-perform-in-eisenstadt")
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
