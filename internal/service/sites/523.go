package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST523(email string) (int, error) {
	urlString := "https://www.petissimo.de/newsletter-abonnieren"

	data := url.Values{}
	data.Set("email_txt", email)
	data.Set("firstname_txt", "dksopfjsxsax")
	data.Set("lastname_txt", "fdasfasdxs")
	data.Set("newsletter_categories_blob[]", "13")
	data.Set("newsletter_categories_blob[]", "14")
	data.Set("newsletter_categories_blob[]", "36")
	data.Set("newsletter_categories_blob[]", "44")
	data.Set("newsletter_categories_blob[]", "48")
	data.Set("ok", "1")
	data.Set("back", "0")
	data.Set("deletion_form", "0")
	data.Set("is_form", "1")
	data.Set("ready_form", "1")
	data.Set("parents", "W10=")
	data.Set("g-recaptcha-response", "03AKH6MRGWoRWR8KKYcTAQEW4tRKjI7HRppa0-cpAU1RbT-pXqSPWQEyFexvA-4FWtnpy8EZU_iGt9cHlolzo84HlvYd0Xsxygph7O2HFQwAzflGZp64PGtd0VCPkP_9SLaXs6I4t-nc6tj6SpQnOQVrUZ8aAHbcVdlbF0I2sQFzQE0kHrkwMK68kfvxstgGXfhxpdUuDJsAxz4o4b8Xds2S_JypK-DlN4GZommVyPv6qRRngvJV-MfmG6EYIrcsoPYfI8Mg-QFLV1U_aThDa1XBg-HyOhJvI7W1GHGW-HEdtVVyR4vOcTUJ3f4m8-Z5oSzjHoGtgOiRhxSFO6H4LvBAkCMW7_TVE8w_Hk69z7Z_c2SjfbRNuzrs2yTsuoCLcHKlqNAMR1Nx_rm3RlW1toYMOLWvtMAhjv8aMvn4X3PxK4_YjeJczdwcDpwS_dSmo22fQ6S80Mu7hMrfFTycy_1Tn4ZoAvvB3yHTU8_azKpxE1PZWOXe1woA60xpwcRluPHhHcQqIfbEHgT2-HQ8zFZv5lGPGr1O396A")

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
