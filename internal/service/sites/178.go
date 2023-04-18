package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST178(email string) (int, error) {
	urlString := "https://www.bsi.bund.de/SiteGlobals/Forms/Newsletter/Newsletter_Bestellen_Formular.html"

	data := url.Values{}
	data.Set("subscriber", email)
	data.Set("idAttributeValue2", "20")
	data.Set("idAttributeValue2.GROUP", "1")
	data.Set("datenschutzhinweis", "Datenschutz")
	data.Set("datenschutzhinweis.GROUP", "1")
	data.Set("submit", "Bestellen")
	data.Set("idAttributeValue1", "173")
	data.Set("idAttributeValue1.HASH", "218248nA3ASad8BCqDatrBRhd14Xihw=")
	data.Set("nn", "520130")
	data.Set("resourceId", "146418")
	data.Set("input_", "520130")
	data.Set("pageLocale", "de")
	data.Set("csrftoken", "683A9C9269DA621F15BA42E788F6D595_630849")

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
