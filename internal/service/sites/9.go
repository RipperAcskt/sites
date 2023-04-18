package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST9(email string) error {
	urlString, _ := url.Parse("https://kcmo.us11.list-manage.com/subscribe/post-json?u=3e4b5887e05b4ab725a02dc40&id=9ae5ee6696&c=jQuery19009713832618894144_1679306750796&EMAIL=123%40gma&b_3e4b5887e05b4ab725a02dc40_9ae5ee6696=&subscribe=Subscribe&_=1679306750797")

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
