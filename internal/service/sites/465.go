package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST465(email string) (int, error) {
	urlString := "https://ui.ungpd.com/Api/Subscriptions/b610bebf-06ee-4ccb-a084-4f1ae911650f"

	data := url.Values{}
	data.Set("Contact[Email]", email)
	data.Set("ConsentText", "I agree to processing of my personal data.")
	data.Set("ListIds", "6e639ecb-6b43-41f6-b837-916b4c08bc45")
	data.Set("SubscriptionConfirmedUrl", "https://www.jpiamr.eu/newsletter-subscription-welcome/")
	data.Set("SubscriptionFailedUrl", "https://www.jpiamr.eu/newsletter-subscription-error/")
	data.Set("DoubleOptIn[Issue][IssueId]", "7d33ccdf-52ff-4da3-8e21-e0e652fd18af")
	data.Set("DoubleOptIn[EmailSentUrl]", "https://www.jpiamr.eu/newsletter-subscription-confirmation/")

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
