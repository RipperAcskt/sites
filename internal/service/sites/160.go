package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST160(email string) (int, error) {
	urlString := "https://rp-online.de/app/newsletter/"

	data := url.Values{}
	data.Set("email", email)
	data.Set("g-recaptcha-response", "03AKH6MRFp574VzGPwcq1ZA1xY60NdQ9HvP2zp_oOrEPRkxEaFggU3UDnpjXMHFNuHJS8W7wDXN3VNN0U4UCt6S3Zyqmuj4V6Fw5VTz-P2xBrHoTDRLwAaYaUHJUMJXcI9GAe9Ov7y7Dgw5y4b_Baky3VR0AN5zlfXKcYFaGYnvB_dMfs_DLdGirLJ9PO5AjDwaEC6XTPuX8s1iqKUxw49mwo8kpW73g5V1pEGOfnbK7yiEfesa3cnskXYCX6_QkwzPxPJ8vx5rUuuz7vnRJIZryBJ969Fyzve4ZGb2ovOObIiwJFRp_ViDDFDYm1-eUiVP4X7qCr4n7YYJQb93IYfeIZ6jiR4nILuEOogFfbKU78n0ljLjG5Dl_kegtD9Br_WcLb72cZozogY1SBb2rqRVUHzXKOGqBEbLfuuwfgkadN_jBiAk3mufvf-j3XaF2ypkGrYKOP__9CCz0VZGGPktmLJqakrOqSFl8gLIAG5top_fY2MoS2Rc2Tbhefp2-MVzEajQcorDmveXqgpurWmr63pGAFgv_kg2Q")
	data.Set("send", "1")
	data.Set("bestellen", "1")
	data.Set("optin", "1")
	data.Set("debug", "1")
	data.Set("anrede", "Herr")
	data.Set("vorname", "qwer")
	data.Set("nachname", "qwer")
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
