package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST113(email string) (int, error) {
	urlString := "https://nyhetsbrev.svt.se/registrering"

	data := url.Values{}
	data.Set("email", email)
	data.Set("terms", "on")
	data.Set("age", "on")
	data.Set("js", "true")
	data.Set("page", "/")
	data.Set("list:62e3ada4f52036c573a42fa3b7264993", "on")
	data.Set("list:c35071e813ee368af6b8e0c091a7d81d", "on")
	data.Set("list:293451b80e8bce63fa56d59f3bf80926", "on")
	data.Set("list:0a373ec1f5b91e59f120f103e020dbc0", "on")
	data.Set("list:c310290dcae0945ba5e20b31ec9ce6dc", "on")

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
