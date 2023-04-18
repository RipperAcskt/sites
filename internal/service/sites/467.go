package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST467(email string) (int, error) {
	urlString := "https://www.esticom.eu/SiteGlobals/Forms/Webs/ESTICOM/newsletter/newsletter-bestellen-formular.html"

	data := url.Values{}
	data.Set("subscriber", email)
	data.Set("pageLocale", "en")
	data.Set("nn", "8804408")
	data.Set("resourceId", "8611898")
	data.Set("input_", "8804408")
	data.Set("submit", "subscribe")
	data.Set("idAttributeValue", "1109")
	data.Set("idAttributeValue.GROUP", "1")

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
