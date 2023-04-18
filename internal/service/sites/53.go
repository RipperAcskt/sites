package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST53(email string) error {
	urlString := "https://popups1-show.getresponse.com/api/popups/ecd6fba2-39ff-440a-a26b-da0ffb6c53eb/contacts"

	data := fmt.Sprintf(`{"campaignId":219418402,"email":"%s","dayOfCycle":0,"customFields":[],"consents":[],"doubleOptIn":false,"pageId":"fc1539de-766b-4cc2-aae8-453860a372ac","formId":"wb-form-7e9f49dcc3d6"}`, email)

	req, err := http.NewRequest("PUT", urlString, strings.NewReader(data))
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
