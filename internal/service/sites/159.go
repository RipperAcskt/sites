package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST159(email string) (int, error) {
	urlString := "https://www.bundesfinanzministerium.de/SiteGlobals/Forms/Newsletter/Newsletter_Bestellen_Formular.html"

	data := url.Values{}
	data.Set("subscriber", email)
	data.Set("privacy.GROUP", "1")
	data.Set("submit", "Absenden")
	data.Set("resourceId", "3aae30a6-d248-4317-a9b6-021da933121b")
	data.Set("input_", "75144d07-41e9-40f4-bad6-36a15833dd5e")
	data.Set("pageLocale", "de")
	data.Set("lists_noTopics.GROUP", "1")
	data.Set("Artikel.GROUP", "1")
	data.Set("Gesetze.GROUP", "1")
	data.Set("Pressemitteilungen.GROUP", "1")
	data.Set("lists_topics.GROUP", "1")
	data.Set("format", "HTML")
	data.Set("format.GROUP", "1")
	data.Set("lists_noTopics", "BMF-Highlights")
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
