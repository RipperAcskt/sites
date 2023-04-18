package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST145(email string) (int, error) {
	urlString := "https://kiddymoon.de/settings.php"

	data := url.Values{}
	data.Set("mailing_email", email)
	data.Set("iai-recaptcha-token", "03AKH6MRG9Hvd_B0-C_m3gZz-5XT1T1-VYukGqVikT1dVzuuQ7yTNHIs4jhm1UPaxw-6G7jqFO4n4pgySCX4_kmfBSHpJhoHUWF0LGmoDL33-xS09GuHMZSUNyIULOkj41qytS4d60LJMvM2vj-R-uqOqJReK-mRlvyOHYQu_s60PDJj6cwyfyuL3D4ai8oAiI0IW_DKDMBwo770MbGEdFcX4mmtk3G1BwQfeZIMYcO8bM-1FP3YyFZutbYpkDgUzCEtlC2kqNOBiey5sFe3fqxJphNy7162QoIV0ybO-vtq7E8ucw9BhHK0BhHwPdLoJSjBEJfO42O90iUuC4_j8ILi6YCtTsCmjr2k66kTMTdW5wuPATmrlvYI1M_KvaOKa_it8x81AEeGmH1_e419ctbpmk6o2bnjPo9aqjo9Nnby98N8SPENzE8ARrb_t2U6bAszFmpJ_x47-Ck1oowLulQwGJvrko2DbtQfmCQzRHTceOX7lpZTiEg-Qw50C4W3vno3xqXVywG-c0")
	data.Set("mailing_action", "add")
	data.Set("newsletter_consent", "on")
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
