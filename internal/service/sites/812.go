package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST812(email string) (int, error) {
	urlString := "https://sumo.com/list/subscribe"

	data := fmt.Sprintf(`{"site_id":"23667734faf53f7c88cf4b96baa776d2a3a8d6865915018b7d037830bf7720b7","visitor_id":"3f856fc458bd0d6afc3a033653af47ee7ae133aa7a4842a62b896a5619afd290","app_id":"156085c5-0017-4150-b225-a731ad248f38","popup_id":"7417b65c2153a595e4f22d93f7f94a940d08f42ba0a9ab336c6969d39d3f4398","object_id":"7417b65c2153a595e4f22d93f7f94a940d08f42ba0a9ab336c6969d39d3f4398,2e37ddf24e696989f4019b98444534b29554427abc310f2e813b0caaeaec29b5","sumolist_id":"2e37ddf24e696989f4019b98444534b29554427abc310f2e813b0caaeaec29b5","href":"https://www.themarginalian.org/","ref":null,"fields":"{\"email\":{\"name\":\"email\",\"value\":\"%s\"},\"name\":{\"name\":\"name\",\"value\":\"qwer\"}}","discount_campaign_id":null,"list_id":null}`, email)

	req, err := http.NewRequest("POST", urlString, strings.NewReader(data))
	if err != nil {
		return 0, fmt.Errorf("new request failed %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("revision", "2022-02-16")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do failed %s", err.Error())
	}
	file, _ := os.Create("POST812")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
