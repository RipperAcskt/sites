package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST506(email string) (int, error) {
	urlString, _ := url.Parse("https://xbrl.us3.list-manage.com/subscribe/post-json?u=fbbc983fbc59adb89ff279b5b&id=da5920711b&c=jQuery190007512515198063041_1680193864393&EMAIL=123dsadsa@gmail.com&FNAME=fasdfasd&LNAME=fdsifosdijf&MMERGE3=fsadfasdf&MMERGE4=fdaskfasd&MMERGE5=dfasjpdfjaps&interestgroup_field=&gdpr[1]=Y&gdpr[5]=Y&b_fbbc983fbc59adb89ff279b5b_da5920711b=&subscribe=Subscribe&_=1680193864394")

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
