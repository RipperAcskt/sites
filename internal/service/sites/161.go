package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST161(email string) (int, error) {
	urlString := "https://www.dgb.de/service/newsletter/"

	data := url.Values{}
	data.Set("555a9a3e-ecb9-11e3-bee2-52540023ef1a_elainenewsletter.email", email)
	data.Set("555a9a3e-ecb9-11e3-bee2-52540023ef1a_elainenewsletter.elaine_selected_newsletters", "b2531e7bb29bf22e1daae486fae3417a")
	data.Set("555a9a3e-ecb9-11e3-bee2-52540023ef1a_elainenewsletter.consent", "on")
	data.Set("555a9a3e-ecb9-11e3-bee2-52540023ef1a_elainenewsletter.actions.unsubscribe", "Abbestellen")


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
