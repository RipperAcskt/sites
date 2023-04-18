package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST940(email string) (int, error) {
	urlString := "https://bbox.blackbaudhosting.com/webforms/components/custom.ashx?handler=blackbaud.appfx.mongo.parts.postformhandler"

	data := url.Values{}
	data.Set("bboxsignup-b0b08373-4cc4-4ded-bf8b-1ae050c78ff9_emailaddress", email)
	data.Set("bboxsignup-b0b08373-4cc4-4ded-bf8b-1ae050c78ff9_firstname", "qwer")
	data.Set("bboxsignup-b0b08373-4cc4-4ded-bf8b-1ae050c78ff9_lastname", "qwer")
	data.Set("bboxsignup-b0b08373-4cc4-4ded-bf8b-1ae050c78ff9$hdnConsentStatement", "Keep up with the latest news from Lupus Canada with our monthly e-newsletters.")
	data.Set("instanceId", "b0b08373-4cc4-4ded-bf8b-1ae050c78ff9")
	data.Set("partId", "f11665b7-d807-48e3-8210-5c47f5766948")
	data.Set("srcUrl", "https://www.lupuscanada.org/news/e-newsletters/")
	data.Set("bboxsignup-b0b08373-4cc4-4ded-bf8b-1ae050c78ff9$btnSubmit", "Submit")
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
