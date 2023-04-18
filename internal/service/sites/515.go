package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST515(email string) (int, error) {
	urlString, _ := url.Parse("https://fyult.maillist-manage.com/weboptin.zc?FIRSTNAME=gdfsgsd&LASTNAME=gsdfgsdf&CONTACT_EMAIL=123asdfg@gmail.com&SIGNUP_SUBMIT_BUTTON=Join Now&zc_trackCode=ZCFORMVIEW&viewFrom=URL_ACTION&submitType=optinCustomView&lD=1ceb23b3129574d5&zx=12d94cd13&zcvers=3.0&mode=OptinCreateView&zcld=1ceb23b3129574d5&zc_formIx=3zc277ebf731a6eb208e7a57309d107df1292d281fef891d1b79015cef920271a7&lf=1680197271596&di=114581921023093825871680197271599&responseMode=inline&sourceURL=https://www.clockwork.io/subscribe/&tpIx=3z75ddb313819b21ba523cafc6b9f8877c8d56a0cc0da9dbffad7d4d91e9c2164a&custIx=3z75ddb313819b21ba523cafc6b9f8877c5a153d657700a3ba10092a306df06f36&cntrIx=3zc277ebf731a6eb208e7a57309d107df1747f9f72f5a579cdbe18fe4a8e5e9d27")

	values := urlString.Query()
	values.Set("CONTACT_EMAIL", email)
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
