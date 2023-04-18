package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST944(email string) (int, error) {
	urlString := "https://viennapride.at/wp-admin/admin-ajax.php"

	data := url.Values{}
	data.Set("field1", email)
	data.Set("field2[]", "zustimmung")
	data.Set("action", "formcraft3_form_submit")
	data.Set("id", "22")
	data.Set("location", "https://viennapride.at/en/")
	data.Set("type", "all")
	data.Set("triggerIntegration", "undefined")
	data.Set("fieldLabels", `{"field2.label":["I consent to the privacy policy."]}`)
	data.Set("formcraft3_wpnonce", "undefined")
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
