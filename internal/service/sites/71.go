package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST71(email string) (int, error) {
	urlString := "https://lp.constantcontactpages.com/su/wveA0ps"

	data := fmt.Sprintf(`{"custom_fields":[],"email_address":"%s","first_name":"qwer","last_name":"qwer","g-recaptcha-response":"03AKH6MRHqmywCzuUnPQQO6MzjnSRlmv7PKWWT8-rqwr6tmzswuiKLTc_ScWGuxs3O_M_YEJNkDXiIRZ12CZxxS5qFcRaeckifBh7Wvk_BrQ1FcqdkbUiY8Qgibub0_P5onSehy7iv5mTCL5wszjnA_OYcND2q0UUZ0V-1lOZvmVeJCARw3PkunsrdyQ48rLv6OAwBu1pq3r_88AjjOZpk5JnA2u547qsy_GHEuEapbsbTz4YaccDrUCKXRPHE0Vh1nVLdu92pMuzoEOg0hf0E-JS2YeuV6R2DEIWUqzA_lFNTXa_RPygWuTQTTB3OxVXh814Aid8AzNSFtVM1AXZPQUO4h7Ib9yTlHlUfyiMtjms-snZwswcv-OjKPhfF9ms9tp9Vwipta2mreHk9Ar7-Cbv7PTN7gdxZoJfP6ahBH94db7Y1EgULl10CbOZuvdd2CgVMFMTO-jkOwULp7V5DqaC2IquxHoGfG9AGJ64Tp83swwdkbu_9YKJhYDJcyJ-2SfcSGckToVNyAj4SL0_604bMiusQXRrRhG60N1zwVD9gqk2cT5ZNHwA"}`, email)

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

	return response.StatusCode, nil
}
