package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST452(email string) (int, error) {
	urlString := "https://www.flexmail.eu/subscribe"

	data := url.Values{}
	data.Set("fields[2426ced136024a631a2a0672a269b700]", email)
	data.Set("public-key", "363282b81069bc0a96cffa8f84e467448e673")
	data.Set("public-fields[117979cc352353c57a691a947b04167f]", "fsadfas")
	data.Set("fields[1a2f3e3b4cebe60c0e9659520e7e121a]", "dfasdfasd")
	data.Set("fields[33ce9f525b80dff17154d826716798ebe0291c5845d2e3349615c72e6c45bb8a44fee17864e79e4fd2fce836eb5fbd3d]", "fasdfa")
	data.Set("fields[a5b5c74ceac36bf3f4144d9f6d53033ef1efda10cc162044115bcdf3cdf3319184e097ca7cb7094670335548f166243a]", "sdfasdf")
	data.Set("preferences[fa6cda94969d54e1104eeaa2a557da90]", "fa6cda94969d54e1104eeaa2a557da90")
	data.Set("preferences[9a5e8aef36a017113316bb49384e9b30]", "9a5e8aef36a017113316bb49384e9b30")
	data.Set("mid[]", "7bcf712f066e391165f3814f7013b48e")

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
