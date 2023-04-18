package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST169(email string) (int, error) {
	urlString := "https://alp.dillingen.de/akademie/akademieberichte/ajaxjson/?tx_providertpl_newsletter%5Baction%5D=api&tx_providertpl_newsletter%5Bcontroller%5D=Newsletter&tx_providertpl_newsletter%5Bformat%5D=json&tx_typoscriptrendering%5Bcontext%5D=%7B%22record%22%3A%22pages_699%22%2C%22path%22%3A%22tt_content.list.20.providertpl_newsletter%22%7D&cHash=aad5e45b41848dbdf35579214098171a"

	data := url.Values{}
	data.Set("email", email)
	data.Set("listId", "news.alp.dillingen.de")
	data.Set("tx_providertpl_newsletter[action]", "api")
	data.Set("tx_providertpl_newsletter[controller]", "Newsletter")
	data.Set("tx_providertpl_newsletter[format]", "json")
	data.Set("tx_typoscriptrendering[context]", `{"record":"pages_699","path":"tt_content.list.20.providertpl_newsletter"}`)
	data.Set("cHash", "aad5e45b41848dbdf35579214098171a")
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
