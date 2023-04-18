package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST17(email string) error {
	urlString := "https://subscriptions.cbc.ca/api/subscription/subscribe"

	data := fmt.Sprintf(`{"newsletter_id":"NEW002","email":"%s","t_i":"2023-03-20T10:16:37.776Z","t_p":"2023-03-20T10:16:51.337Z","hp":"","path":"/","action":"subscriptionCentre"}`, email)

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
