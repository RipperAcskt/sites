package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST462(email string) (int, error) {
	urlString := "https://s338644260.t.eloqua.com/e/f2"

	data := url.Values{}
	data.Set("emailAddress", email)
	data.Set("elqFormName", "SK-Subscriptiontonewsletters")
	data.Set("elqSiteId", "338644260")
	data.Set("singleCheckbox10", "on")
	data.Set("singleCheckbox4", "on")
	data.Set("singleCheckbox2", "on")
	data.Set("singleCheckbox6", "on")
	data.Set("singleCheckbox7", "on")
	data.Set("singleCheckbox9", "on")
	data.Set("singleCheckbox", "on")
	data.Set("firstName", "dsaf")
	data.Set("lastName", "fasdfasd")
	data.Set("company", "dfasdfasd")
	data.Set("title", "fsadfasdf")
	data.Set("phoneNumber1", "12313212")

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
