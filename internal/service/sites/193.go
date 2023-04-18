package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST193(email string) (int, error) {
	urlString := "https://lenta.vn.ua/add-email.html"

	data := url.Values{}
	data.Set("user[email]", email)
	data.Set("redirect", "https%3A%2F%2Flenta.vn.ua%2Fadd-email.html")
	data.Set("redirectunsub", "https%3A%2F%2Flenta.vn.ua%2Fadd-email.html")
	data.Set("option", "com_acymailing")
	data.Set("hiddenlists", "2")
	data.Set("acyformname", "formAcymailing10991")
	data.Set("ajax", "0")
	data.Set("acy_source", "module_97")
	data.Set("ctrl", "sub")
	data.Set("task", "optin")
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
