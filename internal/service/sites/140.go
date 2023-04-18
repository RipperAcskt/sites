package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST140(email string) (int, error) {
	urlString := "https://www.gaggenau.com/de/sys/service/newsletter/subscribe"

	data := url.Values{}
	data.Set("email", email)
	data.Set("optinData", "QTE1X0RFX0QyQy9ORVdTTEVUVEVS")
	data.Set("optinData.consents[0].children[0].children[0].children[0].selected", "true")
	data.Set("_optinData.consents[0].children[0].children[0].children[0].selected", "on")
	data.Set("useSimplifiedForm", "false")
	data.Set("title", "0002")
	data.Set("firstname", "qwer")
	data.Set("lastname", "qwer")
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
