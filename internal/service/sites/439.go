package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST439(email string) (int, error) {
	urlString := "https://api.targito.com/v1.0/contacts/AddContactWeb"

	data := fmt.Sprintf(`{"email":"%s","accountId":"29805a55-787c-43b1-996f-ff32604c0768","origin":"homeandcook_sk","columns":{"gdpr_marketing_parent":"rowenta","gdpr_marketing_event":"welcome","gdpr_marketing_date":"2023-03-28 21:50:52","rowenta_event":"welcome","rowenta_date":"2023-03-28 21:50:52"},"forbidReOptIn":true,"campaignId":"233d3959-fc82-44ff-9b84-729e577cf3c1"}`, email)

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
