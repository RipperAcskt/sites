package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST181(email string) (int, error) {
	urlString := "https://www.bundestag.de/dynamic/newsletteroverviewform/de/services/newsletter/newsletter/854562"

	data := url.Values{}
	data.Set("values['854578']", email)
	data.Set("_csrf", "a5cd9083-466d-4e57-898e-fbe2a5290480")
	data.Set("values['558174'", "true")
	data.Set("_values['558174']", "on")
	data.Set("values['443986']", "artist")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['showAllThemes']", "on")
	data.Set("_values['newsletters']", "on")
	data.Set("_values['newsletters']", "on")

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
