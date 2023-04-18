package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST1(email string) error {
	urlString, _ := url.Parse("https://iecetech.us7.list-manage.com/subscribe/post-json?u=57bf19922547229aeac42e06a&id=00eeae4a79&c=jQuery19007190546441902206_1679305139166&EMAIL=1%40gmail.com&b_57bf19922547229aeac42e06a_00eeae4a79=&subscribe=Subscribe&_=1679305139167")

	values := urlString.Query()
	values.Set("EMAIL", email)
	urlString.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlString.String(), nil)
	if err != nil {
		return fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do failed %s", err.Error())
	}

	fmt.Println(response.StatusCode)
	return nil
}
