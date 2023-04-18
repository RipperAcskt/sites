package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST45(email string) error {
	urlString := "https://api-onsite-v2.edrone.me/onsite-event"

	data := fmt.Sprintf(`{"app_id":"635bc2ecd9b16","action_type":"click","click_data":{"email":"123@gmail.com","first_name":"sg","last_name":"","phone":"","gender":"","customer_tags":"From PopUp","birth_date":"","e_ref":"","action":"subscribe","sms_subscriber_status":1,"post_request_url":"","redirect_url":""},"onsite_type":"popup_v2","onsite_id":15939,"event_date":1679310910598,"publish_version":10,"fpccid":"b5f6f390c89f9742bf299fec8c07d60c82ae73bb","email":"%s","location_url":"https://www.arante.pl/webpage/newsletter.html"}`, email)

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
