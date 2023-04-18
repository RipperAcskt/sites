package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST156(email string) (int, error) {
	urlString := "https://proasyl.gu-marketingsuite.com/proc.php"

	data := url.Values{}
	data.Set("email", email)
	data.Set("act", "sub")
	data.Set("v", "2")
	data.Set("or", "e9761a9b8164668b9fe4043ed821bba0")
	data.Set("field[1]", "Herr")
	data.Set("fullname[1]", "qwer")
	data.Set("u", "1")
	data.Set("f", "1")
	data.Set("c", "0")
	data.Set("m", "0")
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
