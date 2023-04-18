package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST138(email string) (int, error) {
	urlString := "https://www.bundesgerichtshof.de/SiteGlobals/Forms/Newsletter/Newsletter_BGH_Bestellen_Formular.html;jsessionid=1932B4872AC70A579109EAB2C122CFC1.internet941"

	data := url.Values{}
	data.Set("subscriber", email)
	data.Set("idAttributeValue1", "40")
	data.Set("idAttributeValue1.GROUP", "1")
	data.Set("zustimmungdaten", "1")
	data.Set("zustimmungdaten.GROUP", "1")
	data.Set("submit", "Bestellen")
	data.Set("nn", "11363328")
	data.Set("resourceId", "11363364")
	data.Set("input_", "11363328")
	data.Set("pageLocale", "de")
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
