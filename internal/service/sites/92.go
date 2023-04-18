package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST92(email string) (int, error) {
	urlString := "https://collector-pxhd2cxrgh.px-cloud.net/api/v2/collector"

	data := url.Values{}
	data.Set("EMAIL_ADDRESS_", email)
	data.Set("_ri_", "X0Gzc2X%3DAQjkPkSRWQG5uzgiAnzcKr8KP7iXKjvzfyzeyv4IJzdYOVwjpnpgHlpgneHmgJoXX0Gzc2X%3DAQjkPkSRWQGzaD878jeApzgbvRcize9JIRdpkCFRbUO")
	data.Set("_ei_", "E_1LQrRTQ7XzOP9WoZibiwg")
	data.Set("_di_", "r9t3au1tfpq8ceuje92dk3g04c3pn3leei1njgcjlqhu3fpvsgng")
	data.Set("EMAIL_FORMAT_", "H")
	data.Set("FIRST_NAME", "qwer")
	data.Set("LAST_NAME", "qwer")
	data.Set("OB_TITLE", "Marketing")
	data.Set("COMPANY_NAME", "qwer")
	data.Set("POSTAL_STREET_1_", "qwer")
	data.Set("POSTAL_STREET_2_", "qwer")
	data.Set("CITY_", "qwer")
	data.Set("COUNTRY_", "qa")
	data.Set("POSTAL_CODE_", "1234")
	data.Set("MARKET_SEGMENT", "3-X")
	data.Set("TRANSPORTATION", "Y")
	data.Set("lid", "on")


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
