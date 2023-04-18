package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST950(email string) (int, error) {
	urlString := "https://subscriptions.cbc.ca/api/subscription/subscribe"

	data := fmt.Sprintf(`{"newsletter_id":"REG006","email":"%s","t_i":"2023-03-28T22:31:23.199Z","t_p":"2023-03-28T22:31:41.718Z","hp":"","path":"/","action":"subscriptionCentre"}`, email)

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
	file, _ := os.Create("POST950")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
