package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST162(email string) (int, error) {
	urlString := "https://seu2.cleverreach.com/f/290337-324413/wcs/"

	data := url.Values{}
	data.Set("email", email)
	data.Set("1129869[]", "einwilligung")
	data.Set("csrf", "4xYAWwN2AH2g7I17dO9lgRoUbCA")
	data.Set("mgnlModelExecutionUUID", "ad45d79b-0d5b-4d8b-8ff0-b8dfae87493d")
	data.Set("1122810", "qwer")
	data.Set("1122809", "qwer")	
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
