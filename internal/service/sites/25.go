package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST25(email string) error {
	urlString := "https://a.klaviyo.com/client/subscriptions/?company_id=SQhApy"

	data := fmt.Sprintf(`{"data":{"type":"subscription","attributes":{"list_id":"UgLGqi","custom_source":"Newsletter Signup","email":"%s","properties":{"$first_name":"asdf","Gender":"Gender:DoesntID","Birthday":"12/03/2003","$consent_method":"Klaviyo Form","$consent_form_id":"TdkfbY","$consent_form_version":8363514,"services":"{\"shopify\":{\"source\":\"form\"}}","$timezone_offset":3,"$exchange_id":"fMWReeTlIzgNkVYpOgFVLM1rFBAqr0Zq1b8m18BI0Gg=.SQhApy"}}}}`, email)

	req, err := http.NewRequest("POST", urlString, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("revision", "2022-02-16")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do failed %s", err.Error())
	}

	fmt.Println(response.StatusCode)
	return nil
}
