package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST112(email string) (int, error) {
	urlString := "https://ui.ungpd.com/Api/Subscriptions/60a224c8-6f51-4e6b-be09-bc5bab1942db"

	data := url.Values{}
	data.Set("Contact[Email]", email)
	data.Set("SubscriptionConfirmedUrl", "https://ka.se/nyhetsbrev/tack/")
	data.Set("SubscriptionFailedUrl", "https://ka.se/nyhetsbrev/fel/")
	data.Set("DoubleOptIn[Issue][IssueId]", "58342e13-b029-42a0-9e06-064569dded62")
	data.Set("DoubleOptIn[EmailSentUrl]", "https://ka.se/nyhetsbrev/bekrafta/")
	data.Set("ListIds", "7d95911c-f8f5-4582-8bfd-c3957183b26f")
	data.Set("ListIds", "b5859b2a-61df-4888-a1d3-ec20de6e1a83")
	data.Set("ConsentText", "Jag godkänner att LO Mediehus behandlar mina personuppgifter enligt sin personuppgiftspolicy.")

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
