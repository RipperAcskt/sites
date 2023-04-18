package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST192(email string) (int, error) {
	urlString := "https://www.ea.com/api/newsletter/signup"

	data := fmt.Sprintf(`{"token":"SKaG4k2S6mJ7hpIDDtpeBQOihOYX8dHJ6pO5q32bYg1nT2pv8ZJmBoj9S1KZulSMNZ4YyWBJvTOd3LQoP3-1OdlkylY8azXN6wycp5dzxM0ZhURpOaUvDe9toqwI87eONkTfYJpcR_NvSZJQ","email":"%s","dateOfBirth":"1999-12-12","country":"US","preference":""}`, email)

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
