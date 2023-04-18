package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST827(email string) (int, error) {
	urlString := "https://reason.org/wp-admin/admin-ajax.php"

	data := url.Values{}
	data.Set("form_data[0][value]", email)
	data.Set("nonce", "4dc46b988f")
	data.Set("action", "rform_submission_handler_puprf_priority_signup")
	data.Set("form_data[0][name]", "email_email")
	data.Set("form_data[1][name]", "checkboxes_lists[]")
	data.Set("form_data[2][name]", "checkboxes_lists[]")
	data.Set("form_data[3][name]", "checkboxes_lists[]")
	data.Set("form_data[1][value]", "alerts")
	data.Set("form_data[2][value]", "new-at-reason-dot-com")
	data.Set("form_data[3][name]", "checkboxes_lists[]")
	data.Set("form_data[3][value]", "reasonroundup")
	data.Set("form_data[4][name]", "checkboxes_lists[]")
	data.Set("form_data[4][value]", "the-rattler")
	data.Set("form_data[5][name]", "checkboxes_lists[]")
	data.Set("form_data[5][value]", "nyc-events")
	data.Set("form_data[6][name]", "nonce")
	data.Set("form_data[6][value]", "4dc46b988f")
	data.Set("form_data[7][name]", "action")
	data.Set("form_data[7][value]", "rform_submission_handler_puprf_priority_signup")
	data.Set("form_data[8][name]", "form_name")
	data.Set("form_data[8][value]", "4dc46puprf_priority_signupb988f")
	data.Set("form_name", "puprf_priority_signup")

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
