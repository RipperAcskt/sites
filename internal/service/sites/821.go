package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST821(email string) (int, error) {
	urlString := "https://bychefnuit.com/custom_form"

	data := url.Values{}
	data.Set("custom_form_params[Email Address]", email)
	data.Set("custom_form_params[Private Event]", "S Private Eventubscribe")
	data.Set("custom_form_params[form_id]", "1229")
	data.Set("custom_form_params[Name]", "qwer")
	data.Set("custom_form_params[Phone Number]", "1234123412")
	data.Set("custom_form_params[Type of Event]", "Birthday")
	data.Set("custom_form_params[Guests]", "1")
	data.Set("location_id", "760")
	
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
