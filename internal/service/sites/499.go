package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST499(email string) (int, error) {
	urlString := "https://gem.apps.godaddy.com/v1/subscriber"

	data := fmt.Sprintf(`{"accountId":"ffe76b21-a4f8-11ea-81cc-3417ebe724ff","websiteId":"c66a07b1-97c3-43bc-a820-a045f58d13ee","locale":"en-US","formData":[{"label":"Email","value":"%s"}]}`, email)

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
