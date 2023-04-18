package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST945(email string) (int, error) {
	urlString := "https://www.createsend.com/t/securedsubscribe?token=310EC276A7514653F4CB4FE8E399F15882EC1CBEED4891CF9953188801D066DE38635FD2480555A0"

	data := url.Values{}
	data.Set("cm-ejuttl-ejuttl", email)
	data.Set("cm-f-vldljy", "qwer")
	data.Set("cm-f-vldlyu", "qwer")
	data.Set("cm-fo-zhhdur", "3786859")
	data.Set("cm-fo-sjdfy", "3950651")
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
