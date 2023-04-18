package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST108(email string) (int, error) {
	urlString := "https://s.sick.com/e/f2"

	data := url.Values{}
	data.Set("emailAddress", email)
	data.Set("elqFormName", "sys_newsletter_registration")
	data.Set("elqSiteId", "2116941023")
	data.Set("language", "FI")
	data.Set("country", "fi")
	data.Set("DOI_Email", "3340")
	data.Set("tyPage", "2312")
	data.Set("process", "FI-fi")
	data.Set("firstName", "qwer")
	data.Set("lastName", "qwer")
	data.Set("company", "qwer")
	data.Set("optinnewsletter", "1")
	data.Set("optindata", "1")

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
