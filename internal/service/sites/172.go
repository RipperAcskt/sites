package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST172(email string) (int, error) {
	urlString := "https://fsa.gov.ru/press-center/subscribe/subscr_edit.php"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("RUB_ID[]", "2")
	data.Set("RUB_ID[]", "14")
	data.Set("RUB_ID[]", "4")
	data.Set("RUB_ID[]", "6")
	data.Set("RUB_ID[]", "8")
	data.Set("RUB_ID[]", "10")
	data.Set("RUB_ID[]", "12")
	data.Set("RUB_ID[]", "3")
	data.Set("RUB_ID[]", "5")
	data.Set("RUB_ID[]", "7")
	data.Set("RUB_ID[]", "9")
	data.Set("RUB_ID[]", "11")
	data.Set("RUB_ID[]", "13")
	data.Set("TYPE", "d")
	data.Set("FORMAT", "html")
	data.Set("RUB_ID[]", "0")
	data.Set("PostAction", "Add")
	data.Set("sessid", "832a47e2f0a7fd69afe7c73207478860")
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
