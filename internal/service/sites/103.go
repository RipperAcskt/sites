package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST103(email string) (int, error) {
	urlString := "https://www.irac.org/api/form/SaveFormSubmission"

	data := fmt.Sprintf(`{"key":"1:1680132563:mK8NgKrD1SZ4htCF/8zM04x4/ITYS0I0vf2Ii9+7bT0=","formId":"5fd78187047ccb5f522be243","collectionId":"5fd780b1b799144e7f4e3008","objectName":"page-section-5fd78187047ccb5f522be245","form":"{\"name-yui_3_17_2_1_1552579698432_3882\":[\"qwer\",\"qwer\"],\"email-yui_3_17_2_1_1552579698432_3881\":\"%s\"}","pagePermissionTypeValue":1,"pageTitle":"Subscribe to the Pluralist","pageId":"5fd780b1b799144e7f4e3008","contentSource":"c","pagePath":"/sign-up"}`, email)

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
