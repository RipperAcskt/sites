package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST921(email string) (int, error) {
	urlString := "https://news.heckel.fr/frmv2/heckel/FR"

	data := fmt.Sprintf(`{"__RequestVerificationToken":"MTeC_1WFIYNlouVbgIn-fIklKTyR9xmwBz9QzOJ01mr7GvOjxyl3QlQiSYdJjUgciG96R65uGLw1v2EHiQru_lvy37EUJX5kAZ_EVyRdPLwLdNDJN8I50DXt6cayjP2_cdsVHm5r0XZ4JYQOpmTovPco74PM1U8G7A05Lpjxpw41","Salutation":"m","Firstname":"qwer","Lastname":"qwer","Company":"qwer","Industry":"ppe_distributors","Position":"qwer","PostalCode":"123","City":"qwer","Country":"PHL","EMail":"%s","DataProtection":"true","Permission":"true"}`, email)

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
	file, _ := os.Create("POST921")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
