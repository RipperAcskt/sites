package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST968(email string) (int, error) {
	urlString := "https://www.filippo.pl/settings.php"

	data := url.Values{}
	data.Set("mailing_email", email)
	data.Set("mailing_action", "add")
	data.Set("mailing_name", "qwer")
	data.Set("newsletter_consent", "on")
	data.Set("iai-recaptcha-token", "03AKH6MRE8SGwtw93IKVKdiv05-5ESwdyoI6yKwdSK3-6ZWWWvGHj2Der6COZc59FQYIteGEhxeayOY54BKhKLpXzOoRGcG9an8WGjBFiGxzD0DR4b0hdwgXUOU5tKfnqd4zyF1w6bGt3sQwnbF2HehKcaeU4274x2-RrYfpAKdM579ImhOt2xi69QedRR7ByGykbaI4Q3qyetTEYIfO0YDYlrjRKfqKBh4PA2UCFA_Xf9UpCMvTlTeK3D6OMmiXkaXEh8VnfNFCAAMZauK6qUFYIgtdFc7LNKFWhpHvP-gVBwM-L5eunzZVV7-EFlRcKdINLzeGnBTtn06VXklCaESX44RglmyurvDjuzGwzELPZnC2xkC8zOO0kiF6ndzSB4RNv424XJg78O4NgXFdsn4J876aoEAsh5_tfTrCcShvJThZvgpVuawkkEsOoWv-3OtyUfpADE0P56Lt9fquW4o7tVoicGDNwm5npWGEUwRALvzR3yswoiKAQLARGormt2xyz3bp5FkWvD")
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
