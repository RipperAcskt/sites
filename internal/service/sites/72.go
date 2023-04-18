package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST72(email string) (int, error) {
	urlString := "https://gem.apps.godaddy.com/v1/subscriber"

	data := fmt.Sprintf(`{"accountId":"7cec7bc3-82bb-11ea-81c3-3417ebe72601","websiteId":"065b0a3b-c02c-406c-8efb-41fbd105b53e","locale":"en-US","formData":[{"label":"Email","value":"%s"}]}`, email)

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
