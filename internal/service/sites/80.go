package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST80(email string) (int, error) {
	urlString := "https://foodbusinessafrica.us4.list-manage.com/subscribe/post"

	data := url.Values{}
	data.Set("MERGE0", email)
	data.Set("u", "8dacba15942964188af1f3c10")
	data.Set("id", "529878f093")
	data.Set("MERGE1", "qwer")
	data.Set("MERGE6", "qwer")
	data.Set("MERGE3", "qwer")
	data.Set("MERGE2", "Manufacture & Supplies of commodities, processing equipment, post-harvest, ingredients, chemicals, packaging, lab and food safety to the sector")
	data.Set("MERGE4", "Australia")
	data.Set("group[1][1]", "1")
	data.Set("group[1][2]", "1")
	data.Set("group[1][4]", "1")
	data.Set("group[1][8]", "1")
	data.Set("group[1][16]", "1")
	data.Set("group[1][32]", "1")
	data.Set("submit", "Subscribe")
	data.Set("ht", "fb04e78b5811899903e001daa98ad494ea936a42:MTY4MDEyNzkwOS4wOTE1")
	data.Set("mc_signupsource", "hosted")

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
