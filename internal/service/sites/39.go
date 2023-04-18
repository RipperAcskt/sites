package sites

import (
	"fmt"
	"net/http"
	"net/url"
)

func POST39(email string) error {
	urlString, _ := url.Parse("https://web.webformscr.com/members/forms/jsonp-submit?callback=SPFormRegistry[%270db7b7c6dd54bd24fc724213658c749cdbfe518128468a2b014e638f3f80f1aa%27].cbSubmit&sform%5Bemail%5D=123%40gmail.com&=&sform_lang=ua&sform%5Bhash%5D=0db7b7c6dd54bd24fc724213658c749cdbfe518128468a2b014e638f3f80f1aa&sform%5BYXV0b1NpdGU%3D%5D=cambridge.ua&sform%5BYXV0b0lw%5D=185.53.133.75&sform%5BYXV0b0NpdHk%3D%5D=Warsaw&sform%5BYXV0b1JlZ2lvbg%3D%3D%5D=Mazovia&sform%5BYXV0b0NvdW50cnk%3D%5D=Poland&_=1679309808670")

	values := urlString.Query()
	values.Set("sform[email]", email)
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
