package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST458(email string) (int, error) {
	urlString := "https://www.nzcer.org.nz/news/newsletters"

	data := url.Values{}
	data.Set("mergevars[EMAIL]", email)
	data.Set("mergevars[FNAME]", "fasdfasd")
	data.Set("mergevars[LNAME]", "fasdfa")
	data.Set("mergevars[MMERGE3]", "sdfasdfas")
	data.Set("form_build_id", "form-u_9gTAt9aagMOhxGoCkMsIEUtZSNx8dMmfJxmowx0Yw")
	data.Set("form_id", "mailchimp_signup_subscribe_block_nzcer_newsletter_form")
	data.Set("mailchimp_lists[interest_groups][2708c83289][ee7f1fbcb8]", "ee7f1fbcb8")
	data.Set("mailchimp_lists[interest_groups][2708c83289][454c9eff69]", "454c9eff69")
	data.Set("mailchimp_lists[interest_groups][2708c83289][ef6e603191]", "ef6e603191")
	data.Set("mailchimp_lists[interest_groups][2708c83289][1138d69665]", "1138d69665")
	data.Set("mailchimp_lists[interest_groups][2708c83289][07c22a13e7]", "07c22a13e7")
	data.Set("op", "Subscribe")

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
