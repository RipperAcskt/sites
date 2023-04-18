package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST146(email string) (int, error) {
	urlString := "https://www.gtai.de/de/trade/service/newsletter/profil/97676!newsletterSubscribe"

	data := url.Values{}
	data.Set("email", email)
	data.Set("newsletters", "red_gtai_onlinenews")
	data.Set("_newsletters", "on")
	data.Set("salutation", "male")
	data.Set("title", "Prof.")
	data.Set("lastName", "qwer")
	data.Set("_contactAllowed", "on")
	data.Set("contactAllowed", "on")
	data.Set("_csrf", "b5fe420e-4f1f-4f63-b6ed-6b91615d8b1f")
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
