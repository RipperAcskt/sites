package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST498(email string) (int, error) {
	urlString := "https://nutritionfacts.org/subscribe/"

	data := url.Values{}
	data.Set("EMAIL", email)
	data.Set("FNAME", "fspadjfpijas")
	data.Set("INTERESTS[2e381559c0]", "b90f5a65a6")
	data.Set("INTERESTS[a7cbdc5640]", "9d359abe73")
	data.Set("INTERESTS[f178a857e9]", "345ea2bb02")
	data.Set("SIGNUP", "Subscribe to NutritionFacts.org")
	data.Set("_mc4wp_timestamp", "1680190652")
	data.Set("_mc4wp_form_id", "40620")
	data.Set("_mc4wp_form_element_id", "mc4wp-form-1")

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
