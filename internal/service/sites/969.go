package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST969(email string) (int, error) {
	urlString := "https://glam-shop.pl/settings.php"

	data := url.Values{}
	data.Set("mailing_email", email)
	data.Set("mailing_action", "add")
	data.Set("mailing_name", "qwer")
	data.Set("newsletter_consent", "on")
	data.Set("iai-recaptcha-token", "03AKH6MREjy6BxGY_3jFCN5KWUjUVH4PGubCuwKrSRtVMO-pYUf3UUfTFlKzB1QvvaZNAcyR_TCLuqqJc0pvrj7XTasebwKX02Cgk-25u9nUe-cRVvIuDKv9xXrgpnh0WLsbBwOV0pJUQP38hioAPoLFbWtKGgEDtPWUQ0ACvP3l2cWx8psYQs3z1Fhdv7nw_iJRwyCOtMwt9-16JQMr_psHS5rAsyA6mPe4RWEajYFzM77GfqPh4HdJ9f3zYsNgsTB0Ca0OYdjSZqhXPx54jmOb5flw75Wa1kFT4K-EGAhZuLF-Qm6d1SbHKuO2sTUcY4icONJ6KhyaAUmvg7ZtsC6_-1AkWxnJuqa6a22haDdTLBReOIYiVpGbnGrEf9LuR6ErvYa8xh8Pxe6V9NzVprrRixnQ82BFdiJGF8wDl_hRPEJsOrxBLT1QMFJuQXWEVPLtBxaaRpOhSFLxOIp30JF3m01ZaYxMKebRWs4J_pGQ7JhIGRmxto6xIkDjKdwgStEnlkjAMPBU9H")
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
