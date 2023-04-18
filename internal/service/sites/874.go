package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST874(email string) (int, error) {
	urlString := "https://www.createsend.com/t/securedsubscribe?token=41BAC7F5BF3E173D064CE29CF0775F34ACFCD9261AF150C093BE1265442F2F8E5107BD5B3B6992FD"

	data := url.Values{}
	data.Set("email", email)
	data.Set("cm-f-ealiy", "qwer")
	data.Set("cm-f-ealij", "qwer")
	data.Set("cm-fo-ealit", "3912343")
	data.Set("cm-fo-eaxl", "3912394")
	data.Set("cm-privacy-consent", "on")
	data.Set("cm-privacy-consent-hidden", "true")

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
