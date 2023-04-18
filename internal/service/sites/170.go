package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST170(email string) (int, error) {
	urlString := "https://www.nuessing.de/newsletter/addtonewsletter"

	data := url.Values{}
	data.Set("Email", email)
	data.Set("IsOkToNewsletterPrivacy", "true")
	data.Set("IsOkToNewsletterPrivacy", "false")
	data.Set("__RequestVerificationToken", "bbWKRU6cwO1NgYWH9CuU52jsiMUpRM9zt_B32jBLGNBQWdjPhjMWGyM1zVRyPdn_mOBeYvYkULBSpQ6FUPfDMLCQ_M01")
	data.Set("returnUrl", "/service-kontakt/newsletter")
	data.Set("FirstName", "qwer")
	data.Set("LastName", "qwer")
	data.Set("AccountNumber", "1234")
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
