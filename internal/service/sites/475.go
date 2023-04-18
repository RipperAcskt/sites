package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST475(email string) (int, error) {
	urlString := "https://japantoday.com/newsletter"

	data := url.Values{}
	data.Set("email", email)
	data.Set("lang", "en")
	data.Set("g-recaptcha-response", "03AKH6MRE30BXMf0ZLshw5n0v9Kgu1f5vT5h8HcZ1XGPf2zV2FNwbtUBCg1P4VEdve2lzKr__EnFARoXFc80DsHAU4BiuldDcXNGKQ3TwZFkQ62PcUZ9vCD3WoqRppvpiPOeqesmZQVQqrWuTnu16_h-IFdLLlmjIDYRDrjIWRV309WqjsUPWFNQnTD87_PReSk8JZCxICv-jmqaBYPkNic9wmmeicbc2-UHmiq601nBL3GB61sM0NHpjuDxEFHbLfhpyJawnZc2KmpmgeMOvyXgb2fZodaaLbkIIBELKBJSahSM2OzC2EA4vQVkyJlXbQ3vy8Nn9NbxACivuZmvxdUmhA4Dxs5ewyymppFB8ZRdbICFINcHLdaaiYaRPQiTq60YreXECnyHrIn0HT614QseFhHOw4O1dJaHA1KS2mIuAKKQhsxWX3Tj25eDImVEYd6eyBctBg6-qF30O_e7X__fVYtgS-ER8uhYrtKH8dblmH_rSHGASkKagtefVmonAzATgwDY5cEbiD_nXS_EDLy83_w-eUu6tMig")
	data.Set("security", "b8d43a4b78e8de2765f3c387458c0866-82105739019579d21b4cc73640d8d64d")

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
