package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST175(email string) (int, error) {
	urlString := "https://www.createsend.com/t/securedsubscribe?token=55CD63CD585F6E959DB20757E3A0C7C1EAD16EBEE0D93727ECAC6E8AD04C0BF3E85B4AE5E19AB768"

	data := url.Values{}
	data.Set("cm-ourirc-ourirc", email)
	data.Set("cm-name", "qwer")
	data.Set("cm-f-dtdldhtt", "qwer")
	data.Set("cm-fo-dtdldhti", "17886949")
	data.Set("cm-privacy-consent", "on")
	data.Set("cm-privacy-consent-hidden", "true")
	data.Set("token", "55CD63CD585F6E959DB20757E3A0C7C1EAD16EBEE0D93727ECAC6E8AD04C0BF3E85B4AE5E19AB768")

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
