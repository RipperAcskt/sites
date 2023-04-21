package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST919(email string) (int, error) {
	urlString := "https://www.ville-louvres.fr/wp-admin/admin-post.php?action=mailpoet_subscription_form"

	data := url.Values{}
	data.Set("data[form_field_MGJjNGJkZTU0N2M4X2VtYWls]", email)
	data.Set("data[form_id]", "1")
	data.Set("token", "16286416ef")
	data.Set("api_version", "v1")
	data.Set("endpoint", "subscribers")
	data.Set("mailpoet_method", "subscribe")
	data.Set("data[cf_1]", "1")
	data.Set("data[cf_1]", "1")
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