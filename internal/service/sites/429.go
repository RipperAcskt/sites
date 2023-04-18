package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST429(email string) (int, error) {
	urlString := "https://s736811171.t.eloqua.com/e/f2?LP=398"

	data := url.Values{}
	data.Set("C_EmailAddress", email)
	data.Set("elqFormName", "SE_News_Subscribe")
	data.Set("elqSiteId", "736811171")
	data.Set("C_State_Prov", "AL")
	data.Set("C_Country", "AQ")
	data.Set("dEMOSEBusinessClass1", "Equipment Rental/Sales")
	data.Set("dEMOSEJobFunction1", "Food and Beverage Dir/Mgr")
	data.Set("digitaledition", "on")
	data.Set("eventline", "on")
	data.Set("LeadSourceMostRecent", "SE_Newsletter_Sub")

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
