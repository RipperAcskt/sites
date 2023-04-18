package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST489(email string) (int, error) {
	urlString := "https://app.e2ma.net/app2/audience/signup/1836913/1751191/"

	data := url.Values{}
	data.Set("email", email)
	data.Set("member_field_first_name", "fspadjfpijas")
	data.Set("member_field_last_name", "jfdoisjaiofasd")
	data.Set("group_1763479", "1763479")
	data.Set("private_set", "{num_private}")
	data.Set("member_field_company", "fasdfasdfa")
	data.Set("Submit", "Submit")

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
