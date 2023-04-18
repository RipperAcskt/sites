package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST114(email string) (int, error) {
	urlString := "https://www.vardsamverkan.se/RegisterNewsLetterBlock/Submitemailadress?node=1196987"

	data := url.Values{}
	data.Set("EmailAdress", email)
	data.Set("ButtonAction", "Prenumerera")
	data.Set("SuccessMsgforemaildelete", "False")
	data.Set("SuccessMsgforemailsend", "False")
	data.Set("CurrentPageLink", "285661")
	data.Set("CurrentLanguage", "sv")
	data.Set("CurrentBlockLink", "1196987")
	data.Set("ShowMSG", "False")
	data.Set("CurrentBlock.AuthorizationCode", "2JCr4xNlJN4jUG1CMNSLb5F9hOjd5L0i")
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
