package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST167(email string) (int, error) {
	urlString := "https://www.aschaffenburg.de/Aktuelles/Newsletter-/Newsletter-der-Pressestelle/DE_index_4184__receipt_fis~a1db6f9dd8822d716c069fc85f50760d.html"

	data := url.Values{}
	data.Set("Email", email)
	data.Set("kryptisch", "ZCT5NwAAQbYyUtLxIRQro7iUfpiESul+qHgGR9Fgqg815WiPiQ1cUw==")
	data.Set("Nachname", "qwer")
	data.Set("Vorname", "qwer")
	data.Set("Datenschutz", "akzeptiert")
	data.Set("form__submit", "yes")
	data.Set("formular", "send")
	data.Set("iField", "ZCT5NwAAI6EdkjksedEbvh/e9QUihb/B")
	data.Set("Anrede", "bitte auswählen")
	data.Set("Titel", "Dr.")
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
