package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST482(email string) (int, error) {
	urlString := "https://www.westboylston-ma.gov/senior-center/webforms/subscribe-senior-connection-newsletter"

	data := url.Values{}
	data.Set("submitted[email]", email)
	data.Set("submitted[name]", "sfdsfsd")
	data.Set("submitted[address]", "fdsafasdfas")
	data.Set("submitted[town___city]", "dfasdfasdfas")
	data.Set("submitted[state]", "sfdsfsd")
	data.Set("submitted[zip_code]", "123321")
	data.Set("submitted[phone]", "1233234543")
	data.Set("submitted[organization]", "gsdfgsdfgs")
	data.Set("details[page_num]", "1")
	data.Set("details[page_count]", "1")
	data.Set("details[finished]", "0")
	data.Set("form_build_id", "form-7oBPMQUXvwmnPMhOolyE8IZdiJXC576omSiJRdcXB4E")
	data.Set("form_id", "webform_client_form_425")
	data.Set("honeypot_time", "1680175665|fmbO-ZYsI1gkiZRBSvtt6xb06dg7TD6tKumKFVuXonA")
	data.Set("op", "Submit")

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
