package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST119(email string) (int, error) {
	urlString := "https://www.ea.com/api/newsletter/signup"

	data := fmt.Sprintf(`{"token":"9s9FLQR5FKdCDBFYYLw_J2LmiTNmChKq1nkT8zMP12KwkpgXFALhZ6PPOBg6MjKLaafumk_st4ULPGRNoEpkbogW-b83qNW52lXWJg6URnLuWarL-_fstk8XE6HLU4O_7TnT7dUTbnq-hOl-","email":"%s","dateOfBirth":"2002-12-02","country":"AR","preference":""}`, email)

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
