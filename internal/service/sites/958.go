package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST958(email string) (int, error) {
	urlString := "https://www.burgkino.at/newsletter"

	data := url.Values{}
	data.Set("mail[0][value]", email)
	data.Set("form_build_id", "Subscform-aCBdjcBL2I2AlSMDl8OQrMFOHJ629YWs5ymz-J2rC1Eribe")
	data.Set("form_id", "simplenews_subscriptions_block_7dd3721f-04c3-42f0-95d0-477e1bb1ff50")
	data.Set("antibot_key", "Q-GeRDl7j05M3eQuNk_ybd7BSnImiflS93xz9I5Qqes")
	data.Set("op", "Subscribe")

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
