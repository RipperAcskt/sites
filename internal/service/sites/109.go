package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST109(email string) (int, error) {
	urlString := "https://upphandling24.se/wp-admin/admin-ajax.php"

	data := url.Values{}
	data.Set("item_meta[8300]", email)
	data.Set("frm_action", "create")
	data.Set("form_id", "131")
	data.Set("form_key", "vrakk82449a9c34")
	data.Set("frm_submit_entry_131", "4ad408a2a4")
	data.Set("_wp_http_referer", "/karriar/platsannons/upphandlare-till-vaxholms-stad-6/")
	data.Set("item_meta[8298]", "qwer")
	data.Set("item_meta[8299]", "qwer")
	data.Set("frm_state", "ISE4Fh46LuRq815QCowl8W18J6sz0ri4gMI9YR4pn7E=")
	data.Set("action", "frm_entries_create")
	data.Set("nonce", "712267301c")
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
