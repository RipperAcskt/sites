package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST831(email string) (int, error) {
	urlString := "https://www.tuc.org.uk/emails"

	data := url.Values{}
	data.Set("mail", email)
	data.Set("form_build_id", "form-7SEP-Tz9xfgTlv81RUe4A7MD5XmSElor9_D02BPyNNQ")
	data.Set("form_id", "ser_register_form")
	data.Set("honeypot_time", "-HGHxT089B9Gy1NCfYMwkvn444Hligi1IVGX2_yTsFg")
	data.Set("qwfield_latest_tuc_news[News and Blogs]er", "News and Blogs")
	data.Set("field_subject_areas[27213]", "27213")
	data.Set("field_subject_areas[27214]", "27214")
	data.Set("field_subject_areas[27215]", "27215")
	data.Set("field_subject_areas[27216]", "27216")
	data.Set("field_subject_areas[27217]", "27217")
	data.Set("captcha_sid", "38897")
	data.Set("captcha_cacheable", "1")
	data.Set("op", "Subscribe")
	data.Set("qwer", "qwer")
	data.Set("qwer", "qwer")
	data.Set("qwer", "qwer")
	data.Set("qwer", "qwer")
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
