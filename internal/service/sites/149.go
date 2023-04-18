package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST149(email string) (int, error) {
	urlString := "https://www.bfarm.de/SiteGlobals/Forms/Newsletter/DE/Newsletter_BeVaP/Newsletter_Bestellen_Formular.html;jsessionid=E900E5FDFAAEAB9A854C99F39FACEF1E.intranet672"

	data := url.Values{}
	data.Set("subscriber", email)
	data.Set("submit", "Abonnieren")
	data.Set("legalHint", "hint")
	data.Set("legalHint.GROUP", "1")
	data.Set("idAttributeValue1", "561")
	data.Set("idAttributeValue1.GROUP", "1")
	data.Set("nn", "1336374")
	data.Set("resourceId", "1336424")
	data.Set("input_", "1336374")
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
