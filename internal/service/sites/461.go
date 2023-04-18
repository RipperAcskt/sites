package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST461(email string) (int, error) {
	urlString := "https://ui.ungpd.com/Api/Subscriptions/708e10d7-dacd-4e23-9dc7-90d0f7a727ab"

	data := url.Values{}
	data.Set("Contact[Email]", email)
	data.Set("ListIds", "a929d735-506b-48c4-a353-550b807407ce")
	data.Set("SubscriptionConfirmedUrl", "https://www.uu.se/en/news/newsletter/thank-you-newsletter/")
	data.Set("SubscriptionFailedUrl", "https://www.uu.se/en/news/newsletter/something-went-wrong-newsletter/")
	data.Set("DoubleOptIn[Issue][IssueId]", "6be4a85c-436f-44f0-a2c8-c8ef260fb687")
	data.Set("DoubleOptIn[EmailSentUrl]", "https://www.uu.se/en/news/newsletter/confirmationpage-newsletter/")
	data.Set("ConsentText", "I want you to send Uppsala University Newsletter to my inbox.")

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
