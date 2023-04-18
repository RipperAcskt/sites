package sites

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func POST918(email string) (int, error) {
	urlString := "https://ffse.fr/wp-admin/admin-ajax.php"

	data := url.Values{}
	data.Set("param[email]", email)
	data.Set("_wpnonce", "2978d3f66c")
	data.Set("cp-page-url", "https://ffse.fr/sabonner-a-la-newsletter")
	data.Set("https://ffse.fr/sabonner-a-la-newsletter", "cp-uid-e33100f4b1c29c0ce1ebe8fe35d861d6")
	data.Set("param[date]", "28-3-2023")
	data.Set("list_parent_index", "0")
	data.Set("action", "cp_add_subscriber")
	data.Set("list_id", "0")
	data.Set("style_id", "cp_id_bbe18")
	data.Set("msg_wrong_email", "Merci d")
	data.Set("message", "Merci pour votre confiance !")
	data.Set("cp_module_name", "NEWSLETTER FFSE SPORT ENTREPRISE")
	data.Set("cp_module_type", "modal")
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
	file, _ := os.Create("POST918")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
