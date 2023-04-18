package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST98(email string) (int, error) {
	urlString := "https://pc-api.polestar.com/eu-north-1/mulesoft-proxy/"

	data := fmt.Sprintf(`{"operationName":"createLeadRequest","variables":{"body":{"country":"IS","firstName":"qwer","lastName":"qwer","email":"%s","newsletterSubscribed":"2023-03-29T23:12:33.383Z","postalCode":"1234","privacyPolicy":"<p>By submitting, the data provided will be used to perform your request according to the <a href=\"https://www.polestar.com/en-is/legal/privacy/privacy-policy/#marketing-via-e-mail-and-websites\" title=\"Information notice\" target=\"_blank\" rel=\"noopener\">Information notice</a></p>\n<p></p>","source":"Web","campaignSourceCode":"NL subscribed polestar","leadRecordType":"B2C","consentType":"Newsletter Subscription","consentName":"Newsletter Subscription","consentDate":"2023-03-29","preferredLanguage":"en_GB"}},"query":"mutation createLeadRequest($body: LeadCaseInput!) {\n  createLeadRequest(body: $body) {\n    statusCode\n    errorMessage\n    mulesoftErrors\n    __typename\n  }\n}\n"}`, email)

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
