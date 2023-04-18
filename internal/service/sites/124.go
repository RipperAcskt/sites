package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST124(email string) (int, error) {
	urlString := "https://rka.nu/sitevision/proxy/radetforframjandeavkommunalaanalyser/kolada/nyheterikolada.44794.html/svid12_1f376ad3177c89481f780d/-309501482/public/process-subscription-form.aspx;jsessionid=71DDED519BE4D71C9A345E4D5506F126?formId=4645584670454B5D4671"

	data := url.Values{}
	data.Set("pf_Email", email)
	data.Set("pf_CounterMailinglists", "1")
	data.Set("pf_AccountId", "16476")
	data.Set("pf_ListById", "1")
	data.Set("pf_Version", "2")
	data.Set("pf_MailinglistName1", "753715")
	data.Set("Submit", "Prenumerera")
	data.Set("pf_FormType", "OptInForm")
	data.Set("pf_OptInMethod", "SingleOptInMethod")
	data.Set("pf_CounterDemogrFields", "0")
	data.Set("formId", "4645584670454B5D4671")
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
