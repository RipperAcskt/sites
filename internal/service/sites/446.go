package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST446(email string) (int, error) {
	urlString := "https://www.mexicanist.com/members/api/send-magic-link/"

	data := fmt.Sprintf(`{"name":"gsdfgsdf","email":"%s","newsletters":[{"id":"627106b1d597e50031d68d39"},{"id":"64149c310ad152003d3d327f"},{"id":"641a5d8a3c8187003d048eb6"}],"emailType":"signup","requestSrc":"portal","urlHistory":[{"path":"/tag/newsletter/","time":1680042285226,"referrerSource":null,"referrerMedium":null,"referrerUrl":null}]}`, email)

	req, err := http.NewRequest("POST", urlString, strings.NewReader(data))
	if err != nil {
		return 0, fmt.Errorf("new request failed %s", err.Error())
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
		return 0, fmt.Errorf("do failed %s", err.Error())
	}

	return response.StatusCode, nil
}
