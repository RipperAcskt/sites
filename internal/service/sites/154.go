package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST154(email string) (int, error) {
	urlString := "https://www.cduhessen.de/service/newsletter/"

	data := url.Values{}
	data.Set("newsletterabo_value", email)
	data.Set("newsletterabo_list1", "1")
	data.Set("newsletterabo_privacy", "1")
	data.Set("addNewsletterAbo", "Newsletter abonnieren")
	data.Set("unique", "6424ef4147839")
	data.Set("newsletterabo_lastname", "qwer")
	data.Set("newsletterabo_firstname", "qwer")
	data.Set("newsletterabo_salutation", "1")
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
