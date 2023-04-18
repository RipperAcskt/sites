package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST90(email string) (int, error) {
	urlString := "https://www.ihe-europe.net/newsletter/subscribe"

	data := url.Values{}
	data.Set("ergevars[EMAIL]", email)
	data.Set("mergevars[COMPANY]", "qwer")
	data.Set("mergevars[POSITION]", "qwer")
	data.Set("mergevars[LNAME]", "qwer")
	data.Set("mergevars[FNAME]", "qwer")
	data.Set("mergevars[COSTITUENC]", "User/Provider")
	data.Set("mergevars[OTHER]", "qwer")
	data.Set("mergevars[CSIZE]", "2-50")
	data.Set("mergevars[PHONE]", "2341")
	data.Set("mergevars[MOBILE]", "1234")
	data.Set("mergevars[COUNTRY]", "qwer")
	data.Set("mergevars[MMERGE11]", "qwer")
	data.Set("mergevars[MMERGE12]", "qwer")
	data.Set("mergevars[MMERGE13]", "1234")
	data.Set("mergevars[MMERGE14]", "qwer")
	data.Set("mergevars[MMERGE16]", "First Choice")
	data.Set("form_build_id", "form-vu2LIioZoIjv7SsXrDslNaQuenCgWc7AQE8Oh6oFbJA")
	data.Set("form_id", "mailchimp_signup_subscribe_block_block_subscribe_form")
	data.Set("op", "Submit")
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
