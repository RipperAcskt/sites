package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST453(email string) (int, error) {
	urlString := "https://www.ait.ac.at/fileadmin/newsletterAnmeldung/saveNewsletterData.php"

	data := url.Values{}
	data.Set("Email", email)
	data.Set("Emailb", email)
	data.Set("lang", "en")
	data.Set("country", "AT")
	data.Set("anrede", "Ms")
	data.Set("Titel", "fasdfasd")
	data.Set("Firma_Organisation_Institution", "fasdfasd")
	data.Set("Funktion", "fsafdsaf")
	data.Set("Vorname", "dsafsadfas")
	data.Set("Nachname", "dfasdfas")

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
