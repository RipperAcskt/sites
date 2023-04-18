package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST836(email string) (int, error) {
	urlString := "https://austinkleon.substack.com/api/v1/free"

	data := fmt.Sprintf(`{"first_url":"https://austinkleon.substack.com/embed","first_referrer":"https://austinkleon.com/","current_url":"https://austinkleon.substack.com/embed","current_referrer":"https://austinkleon.com/","referral_code":"","source":"embed","referring_pub_id":"","additional_referring_pub_ids":"","email":"%s"}`, email)

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
	file, _ := os.Create("POST836")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
