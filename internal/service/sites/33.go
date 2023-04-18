package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST33(email string) error {
	urlString := "https://blog.craftsmancreative.co/members/api/send-magic-link/"

	data := fmt.Sprintf(`{"email":"%s","labels":[],"autoRedirect":true,"urlHistory":[{"path":"/","time":1679309222101,"referrerSource":null,"referrerMedium":null,"referrerUrl":null}]}`, email)

	req, err := http.NewRequest("POST", urlString, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("revision", "2022-02-16")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do failed %s", err.Error())
	}

	fmt.Println(response.StatusCode)
	return nil
}
