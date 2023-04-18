package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST144(email string) (int, error) {
	urlString := "https://meine.zeit.de/api/1/users"

	data := fmt.Sprintf(`{"email":"%s","login_url":"https://meine.zeit.de/passwort_bestaetigen/{}?entry_service=newsletter-signup-was-jetzt&product_id=sonstige&url=https%3A%2F%2Fwww.zeit.de%2Fnewsletter%2Findex%23newsletter-signup-was-jetzt&wt_zmc=fix.int.zonaudev.zeitde.nl_signup.was-jetzt.nl_signup.link.sonstige&utm_medium=fix&utm_source=zeitde_zonaudev_int&utm_campaign=nl_signup&utm_content=was-jetzt_nl_signup_link_sonstige","short_url":"https://meine.zeit.de/link/{}"}`, email)

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
