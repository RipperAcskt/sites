package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST838(email string) (int, error) {
	urlString := "https://www.aweber.com/scripts/addlead.pl"

	data := url.Values{}
	data.Set("email", email)
	data.Set("submit", " Safe Subscribe")
	data.Set("custom WhereHeard", "qwer")
	data.Set("name", "qwer")
	data.Set("meta_required", "email")
	data.Set("meta_message", "1")
	data.Set("meta_adtracking", "True_Sidebar")
	data.Set("meta_redirect_onlist", "https://thisistrue.com/duplicate/")
	data.Set("redirect", "https://thisistrue.com/success/")
	data.Set("listname", "thisistrue")
	data.Set("meta_web_form_id", "1245921654")
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
