package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST432(email string) (int, error) {
	urlString := "https://www.weidmueller.com/wmcw-servlets/formvalidation/validate?_gl=1*9dzfnd*_ga*MTIxNjk5OTUzNi4xNjgwMDIyMzU5*_ga_WZ2B2RXV5X*MTY4MDAyMjM1OS4xLjEuMTY4MDAyMjM3Mi4wLjAuMA.."

	data := url.Values{}
	data.Set("c_email", email)
	data.Set("formtype", "subscribe")
	data.Set("type", "opt-in")
	data.Set("charset", "utf-8")
	data.Set("okpage", "https://www.weidmueller.hu/hu/egyeb/feliratkozas_megerositese.jsp")
	data.Set("errpage", "https://www.weidmueller.hu/hu/egyeb/hiba_az_oldalon.jsp")
	data.Set("client", "10")
	data.Set("group", "1254")
	data.Set("c_accepted_newsletter", "1")

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
