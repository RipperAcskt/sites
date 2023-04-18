package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST158(email string) (int, error) {
	urlString := "https://www.bibb.de/de/17489.php/subscribe"

	data := url.Values{}
	data.Set("subscribe[email]", email)
	data.Set("subscribe[_token]", "hB0ZKkdl0yckj1cLPECZNuNXkXVF184WMU7Gas4Yvdk")
	data.Set("subscribe[distributor_ids][]", "bf2d84d4-54ac-4879-b112-03b75cdbcd39")
	data.Set("subscribe[privacy_policy]", "1")
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
