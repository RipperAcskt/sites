package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST47(email string) error {
	urlString := "https://dolnyslask.witd.gov.pl/cms/public/request.php"

	data := url.Values{}
	data.Set("newsletterFormEmail", email)
	data.Set("newsletterFormGroup", "1")
	data.Set("newsletterFormAgreement", "1")
	data.Set("Action", "ADD_EMAIL")
	data.Set("PluginName", "newsletter")
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
