package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST925(email string) (int, error) {
	urlString := "https://marketing.cnav.fr/users/subscribe/js_id/6rpe/id/7"

	data := url.Values{}
	data.Set("email", email)
	data.Set("js_id", "6rpe")
	data.Set("from_url", "yes")
	data.Set("hdn_email_txt", "1680013745766")
	data.Set("req_hid", "~NOM~PRENOM~NOM_ASSOCIATION~ORGANISME")
	data.Set("NOM", "qwer")
	data.Set("ORGANISME", "qwer")
	data.Set("PRENOM", "qwer")
	data.Set("NOM", "qwer")
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
