package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST939(email string) (int, error) {
	urlString := "https://org-5beb88cb-3376-4b82-8f73-bb18be1ed624.salsalabs.org/api/activity/submission/subscription"

	data := fmt.Sprintf(`{"header":{},"payload":{"userInteractionId":"d9e382c7-4705-465e-b1d3-59e48f3dd3de","pid":"30690d6f-e02e-4198-b843-77e70de490cd","oid":"5beb88cb-3376-4b82-8f73-bb18be1ed624","aid":"e65e6f06-cc47-4179-b376-2a3220723050","cid":"","eid":"eeec55cf-a628-4acf-bb78-79f064ab4572","eType":"Template","commSrcId":"","commSrcType":"","data":{"PersonCensus@FirstName":{"value":"qwer","label":"First name / Prénom","required":true},"PersonCensus@LastName":{"value":"qwer","label":"Surname / Nom de famille","required":true},"PersonContact@Email@Value":{"value":"%s","label":"Email / Courriel","required":true},"CustomFieldValue@69f7fb13-1b5b-4249-b98c-8e8e1649aac0@VALUE":{"value":"English","label":"Preferred Language / Langue préférée","required":true},"termsAndConditions":{"value":true},"nonce":{"value":"GLmLkEg4daoAICY1g/4qEFekaQPVPKgr/w8XosYJIe89mSkSsqQkTJjyCctlukqWcNVwaP9E9RiJi8m2nmPwvQZCkxTWFO03SZqf0JkN1lhZjm8Vr1ST+DoqaQpLkVh7NXeqm85WpKYDtKEOHe8yHYl3A/HR+qOJZBNj6DQ/rp/nhGKcZjJoNHLJf2r7Ql1G0Z6VyBfqmDDMr/ejXoUrNv9EamC7f4jhgWuXQCwCsAw2EXDbFcSuQFQWBSMQ5f2CCsaWu82XRcHvYG6OfoctW2GMzcKymNjbigSgCvVaaO8MOZTRRn6VauxeRWKVfNcfgncEKCqOJ9wbYNV76rGGlg=="}},"contentChannels":{"9c6838c7-d084-4392-8146-bbf271c912c6":{"label":"Fundraising","optedIn":"true"},"fd7c5589-4046-4ee7-b1ae-b404c8cb5fbf":{"label":"General","optedIn":"true"}},"contactMethods":{"Email":{"label":"Email","optedIn":true,"onForm":false}},"mailingLists":{},"urlMailingLists":"","activityId":"b772be15-3d75-4f3b-91d1-d015d61cd4df","salsaTrack":"","rawForm":{"field-person-firstname":"qwer","field-person-lastname":"qwer","field-contact-email":"qwer@gmail.com","field-69f7fb13-1b5b-4249-b98c-8e8e1649aac0":"English","termsAndConditions":"accepted"},"siftSession":"_igMvBAKVCZOx","showEmailConfirmation":false}}`, email)

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
	file, _ := os.Create("POST939")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
