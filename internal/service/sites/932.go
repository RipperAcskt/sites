package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST932(email string) (int, error) {
	urlString := "https://receiver.emkt.dinamize.com/in/304536/6/05fca/0/?cmp6=resumo-do-dia&cmp33=&cmp45=&cmp34=&cmp112=&cmp114=&cmp40=&cmp18=&cmp35=&cmp10=&cmp42=&cmp30=&cmp48=&cmp23=&cmp94=&cmp39=&cmp24=&cmp19=&update_mode=AV&text-confirmation=U2V1IGUtbWFpbCBmb2kgY2FkYXN0cmFkbyBjb20gc3VjZXNzbyE=&cmp1=qwer@gmail.com&isMsg=true&phase-change=off&form-code=10"

	data := url.Values{}
	data.Set("cmp1", email)
	data.Set("update_mode", "AV")
	data.Set("text-confirmation", "U2V1IGUtbWFpbCBmb2kgY2FkYXN0cmFkbyBjb20gc3VjZXNzbyE=")
	data.Set("isMsg", "true")
	data.Set("phase-change", "off")
	data.Set("form-code", "10")
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
