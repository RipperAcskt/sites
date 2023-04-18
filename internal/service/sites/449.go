package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST449(email string) (int, error) {
	urlString := "https://www.cntraveller.com/newsletters/submit"

	data := url.Values{}
	data.Set("email", email)
	data.Set("chk_traveller", "y")
	data.Set("chk_traveller-postcards", "y")
	data.Set("chk_traveller-restaurants:", "y")
	data.Set("chk_traveller-where-to-go-now", "y")
	data.Set("chk_traveller-travel-better", "y")
	data.Set("source", "newsletter_page")
	data.Set("optin_condenast", "y")
	data.Set("optin_thirdparty", "y")

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
