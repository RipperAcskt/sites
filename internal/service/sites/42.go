package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST42(email string) error {
	urlString := "https://app2.salesmanago.pl/form/q040obg8z3g7rn1k/contact.htm"

	data := url.Values{}
	data.Set("sm-form-email", email)
	data.Set("uuid", "186feb4c7cc-0ec406dc3f12-9e5d29c4-4fe2bb3b-1017016d-5a32906870d7")
	data.Set("validationToken", "ea39cfa44243474a96522e423a8093b1")
	data.Set("sm-form-submit", "Zapisz się")
	encoded := data.Encode()
	req, err := http.NewRequest("POST", urlString, strings.NewReader(encoded))
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do failed %s", err.Error())
	}

	fmt.Println(response.StatusCode)
	return nil
}
