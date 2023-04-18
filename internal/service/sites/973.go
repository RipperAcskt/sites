package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST973(email string) (int, error) {
	urlString := "https://reagle.pl/settings.php"

	data := url.Values{}
	data.Set("mailing_email", email)
	data.Set("mailing_name", "qwer")
	data.Set("mailing_action", "add")
	data.Set("newsletter_consent", "on")
	data.Set("iai-recaptcha-token", "03AKH6MREoYnWgoylaFerDRMbBN3B5Gvn84awJqZGQSJvppxaG2OEJTlZsLnqTnXZkGa7O93V_3B3u-Nn2h5sJRdCa22rJZuRS2nkumGWUemg9DNhSmdHnnIpVtt2UYz0fIdri2jTZb6O5ubdWPwdc_CmzNCaFu-remUvkztFHm2P-IhTjuhZ8rv_OoWRkcIHcIpxSHxnIXhwhZQWzVvFvoWiUHZ7q4M18OU8ZpGipX03xAPIXXGMLJhBqHG6S3MbH5w5sd2L52z0CL2I3YWD8sWmwwkttKOUeWChoH3MzSq_IMSSBws8OudYTABw0D4IvvwFV5j0MZI6x63w8zl-jNuFmUGicdlszE897Ck0WNqSeFu7NE3mz4KYOywqVYcLcpOW5P-6IvaMizIJ6CfbCdEdp0g_X3wZkRNBz1WXquiMMLdkU81p2ynLzpRK-tGmQwpORw37vCggJF8kE_iBQmi2YLt5gPpd9_WYAH4PAcyyz4xcoQJm9fWsYEGGiLsU6mOJKU9kZ6OYM")
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
