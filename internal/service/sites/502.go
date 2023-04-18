package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST502(email string) (int, error) {
	urlString, _ := url.Parse("https://bccie.us2.list-manage.com/subscribe/post-json?u=e771c2afa220cbe24df573c38&id=1c882e8140&c=jQuery1900685459335766142_1680192046811&EMAIL=123dsadsa@gmail.com&FNAME=fasdfas&COMPANY=fasdfas&LNAME=fdsafdsa&MMERGE8=Public+university/college/institute&SECTROTHER=fdsafsda&POSITION=fsadfsadfds&CONSENT=Yes&group[17929][1]=1&group[17929][2]=2&group[17929][8]=8&group[17937][16]=16&b_e771c2afa220cbe24df573c38_1c882e8140=&subscribe=Subscribe&_=1680192046812")

	values := urlString.Query()
	values.Set("EMAIL", email)
	urlString.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlString.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do failed %s", err.Error())
	}

	return response.StatusCode, nil
}
