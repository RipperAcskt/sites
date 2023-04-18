package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST84(email string) (int, error) {
	urlString := "https://app.mailjet.com/widget/iframe/7DF3/N0F"

	data := url.Values{}
	data.Set("w-field-field-188397-1162683-1820757-email", email)
	data.Set("w-field-field-188397-1162683-1820757-12119", "qwer")
	data.Set("w-field-field-188397-1162683-1820757-12120", "qwer")
	data.Set("w-field-field-188397-1162683-1820757-12185", "System Integrator")
	data.Set("w-field-field-188397-1162683-1820757-12122", "Transport")
	data.Set("w-field-field-188397-1162683-1820757-38428", "true")
	data.Set("w-field-field-188397-1162683-1820757-38428", "true")
	data.Set("w-field-field-188397-1162683-1820757-38427", "true")
	data.Set("w-field-field-188397-1162683-1820757-38427", "true")
	data.Set("w-preview-consent-checkbox", "on")
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
