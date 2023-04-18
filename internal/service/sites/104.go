package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST104(email string) (int, error) {
	urlString := "https://www.emp-online.com/on/demandware.store/Sites-GLB-Site/en/EmarsysNewsletter-DialogSubscription"

	data := url.Values{}
	data.Set("email", email)
	data.Set("firstname", "qwer")
	data.Set("lastname", "qwer")
	data.Set("newsletterResource", "1")
	data.Set("csrf_token", "GDwuzi2hR_LA4EPzi6kzyNkuczEv4adPYkTmywlG0RhbJb1tHHMIad-t2HO45i4xxYhuTYc-y9ksSvuR-M33zGjxpIr8FgPjUM5zBHaCjGVW1oQKGbFj54IqiE-7Hk6aupy4wgKHt3CoPQz7qy4r84Xq64hAW43p8lAE_FaeH0iz3b7iedc=")
	data.Set("gender", "1")


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
