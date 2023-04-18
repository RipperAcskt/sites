package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST928(email string) (int, error) {
	urlString := "https://www.heremagazine.ca/newsletter-signup/"

	data := url.Values{}
	data.Set("input_12", email)
	data.Set("input_11.1", "Here Magazine Community News")
	data.Set("is_submit_8", "1")
	data.Set("gform_submit", "8")
	data.Set("state_8", "WyJbXSIsIjU4MTlmNjQzNTBjOTg3ZTFlNWM3YTg5NWRlMjVkMjM4Il0=")
	data.Set("gform_target_page_number_8", "0")
	data.Set("gform_source_page_number_8", "1")
	data.Set("version_hash", "408065c754394e98005899f5249e681d")
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
