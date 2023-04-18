package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST518(email string) (int, error) {
	urlString := "https://www.gff.ge/en/media/subscribe?email=faceitqutick@gmail.com"

	data := url.Values{}
	data.Set("email", email)
	data.Set("first_name", "fspadjfpijas")
	data.Set("last_name", "jfdoisjaiofasd")
	data.Set("birthday[year]", "2000")
	data.Set("birthday[month]", "8")
	data.Set("birthday[day]", "14")
	data.Set("form_build_id", "form-MdKpLr2g8HcpLpNzMg2IIAR7oSfnqQqLGgc215GFYtw")
	data.Set("form_id", "webform_submission_subscribe_node_16260_add_form")
	data.Set("op", "Submit")

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
