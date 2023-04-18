package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST28(email string) error {
	urlString := "https://faisdkapi.mailmodo.com/sdk/event/track"

	data := fmt.Sprintf(`{"user_id":"e9d4f079-54ac-4919-98c2-c31141eff821","event_name":"$form_submitted","event_properties":{"$email":"%s","$page_title":"Email Marketing & Growth Newsletter - Roundup, Best Practices, Industry Insights | Mailmodo","$referrer":"https://www.mailmodo.com/newsletter/","$referrer_domain":"www.mailmodo.com","$referrer_url":"www.mailmodo.com/newsletter/","$page_raw_url":"https://www.mailmodo.com/newsletter/","$page_domain":"www.mailmodo.com","$page_url":"www.mailmodo.com/newsletter/"},"user_properties":{"$platform":"web","$screen_width":1440,"$screen_height":900},"auto":false}`, email)

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
