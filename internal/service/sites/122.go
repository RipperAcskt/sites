package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST122(email string) (int, error) {
	urlString := "http://www.anpdm.com/public/process-subscription-form.aspx?formId=4144594277454B59407240"

	data := url.Values{}
	data.Set("pf_Email", email)
	data.Set("pf_CounterMailinglists", "1")
	data.Set("pf_AccountId", "10188")
	data.Set("pf_ListById", "1")
	data.Set("pf_Version", "2")
	data.Set("Submit", "Prenumerera")
	data.Set("pf_FormType", "OptOutList")
	data.Set("pf_OptInMethod", "SingleOptInMethod")
	data.Set("pf_CounterDemogrFields", "0")
	data.Set("pf_SubscriberName", "qwer")
	data.Set("pf_MailinglistName1", "577157")
	data.Set("pf_DeliveryFormat", "HTML")
	data.Set("formId", "4144594277454B59407240")
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
