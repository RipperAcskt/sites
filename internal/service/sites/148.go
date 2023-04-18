package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST148(email string) (int, error) {
	urlString := "https://mein.shz.de/newsletter/chefredaktionsnewsletter"

	data := url.Values{}
	data.Set("Newsletter[email]", email)
	data.Set("Newsletter[keys][]", "CR-NL")
	data.Set("ovs_csrf", "ENsWgKE6Ck4wHugqoLsIhast3Ku1-O8pwYBEp85TQSR_n3_P6W8_fwl2mlna8XrGmlqfysO0jUjw5yH_qwQuHA==")

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
