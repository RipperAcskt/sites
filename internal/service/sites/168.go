package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST168(email string) (int, error) {
	urlString := "https://www.aschaffenburg.de/Aktuelles/Newsletter-/Newsletter-der-Pressestelle/DE_index_4184__receipt_fis~82911769209bbd1ead4ba6404d3b72c9.html"

	data := url.Values{}
	data.Set("Email", email)
	data.Set("Anrede", "bitte auswählen")
	data.Set("Titel", "Dr.")
	data.Set("Nachname", "qwer")
	data.Set("Vorname", "qwer")
	data.Set("Datenschutz", "akzeptiert")
	data.Set("form__submit", "yes")
	data.Set("formular", "send")
	data.Set("iField", "ZCWDIgANpy+ELo1f/ehakRIQYHgG3c43")
	data.Set("kryptisch", "ZCWDIgANsK6QDKp19xVWyAhf2oezFSH6uNvYJJ/nkrkzzRA5nKvqCQ==")
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
