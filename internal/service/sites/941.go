package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST941(email string) (int, error) {
	urlString := "https://services.international.gc.ca/mc-cm/services/subscribe.aspx"

	data := url.Values{}
	data.Set("email", email)
	data.Set("list", "001_foreign_affairs_eng")
	data.Set("list", "008_international_development_prod_list_fra")
	data.Set("lists", "001_foreign_affairs_eng")
	data.Set("url", "https://www.international.gc.ca/global-affairs-affaires-mondiales/news-nouvelles/email-subscribe-confirm-courriel-abonnement-confirmer.aspx?lang=eng")
	data.Set("secx", "a82931ba")
	data.Set("subscribe", "Subscribe")
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
