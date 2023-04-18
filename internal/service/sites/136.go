package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST136(email string) (int, error) {
	urlString := "https://web.inxmail.com/adesso/subscription/servlet"

	data := url.Values{}
	data.Set("email", email)
	data.Set("Kampagne", "unbekannt")
	data.Set("Nachname", "qwer")
	data.Set("Vorname", "qwer")
	data.Set("Firma", "qwer")
	data.Set("INXMAIL_HTTP_REDIRECT_ABMELDEN", "https://www.adesso.de/de/newsletter-blog/abmeldung.jsp")
	data.Set("INXMAIL_SUBSCRIPTION", "Blog Newsletter")
	data.Set("INXMAIL_HTTP_REDIRECT", "https://www.adesso.de/de/newsletter-blog/anmeldung-erfolgreich-uebermittelt.jsp")
	data.Set("INXMAIL_HTTP_REDIRECT_ERROR", "https://www.adesso.de/de/newsletter-blog/fehler-bei-der-anmeldung.jsp")
	data.Set("INXMAIL_HTTP_REDIRECT_ABGELAUFEN", "https://www.adesso.de/de/newsletter-blog/anmeldungsfrist-abgelaufen.jsp")
	data.Set("INXMAIL_HTTP_REDIRECT_ABGESCHLOSSEN", "https://www.adesso.de/de/newsletter-blog/anmeldung-abgeschlossen.jsp")
	data.Set("INXMAIL_CHARSET", "UTF-8")
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
