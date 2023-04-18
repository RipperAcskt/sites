package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST459(email string) (int, error) {
	urlString := "https://www.anpdm.com/public/process-subscription-form.aspx?formId=41465A4A7143415E4B7740"

	data := url.Values{}
	data.Set("pf_Email", email)
	data.Set("Submit", "Subscribe")
	data.Set("pf_MailinglistName3", "99706")
	data.Set("pf_MailinglistName2", "99709")
	data.Set("pf_MailinglistName1", "101065")
	data.Set("pf_FormType", "OptInForm")
	data.Set("pf_OptInMethod", "SingleOptInMethod")
	data.Set("pf_CounterDemogrFields", "0")
	data.Set("pf_CounterMailinglists", "3")
	data.Set("pf_AccountId", "7529")
	data.Set("pf_ListById", "1")
	data.Set("pf_Version", "2")
	data.Set("pf_redir", "https://forte.se/en/newsletter/?newsletter_site=true")
	data.Set("newsletter-policy", "on")

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
