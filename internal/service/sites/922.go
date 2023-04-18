package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST922(email string) (int, error) {
	urlString := "https://elvstromsails.com/fr/abonnez-vous-a-notre-newsletter"

	data := url.Values{}
	data.Set("form[Email address]", email)
	data.Set("form[First Name]", "qwer")
	data.Set("form[Last Name]", "qwer")
	data.Set("form[Country]", "qwer")
	data.Set("form[Elvstroem Sails Global News][]", "Elvstrøm Sails France Newsletter")
	data.Set("form[formId]", "21")
	data.Set("00c1afd9affff981d73171199eb55fdb", "1")
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
