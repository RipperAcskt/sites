package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST876(email string) (int, error) {
	urlString := "https://registerdisney.go.com/jgc/v8/client/DTCI-NATGEO.CONSUMER.WEB-PROD/marketing?langPref=en-US&feature=no-password-reuse"

	data := fmt.Sprintf(`{"profile":{"email":"%s","countryCodeDetected":"BY","region":"BY"},"legalAssertions":["gtou_ppv2_proxy"],"marketing":[{"code":"WDIGFamilySites","subscribed":true,"textId":"BU_*_LOCATION_AU_bloc__ff6e04dba16ac5ef1d2e28dc22f280fc"},{"code":"National_Geographic_L","subscribed":true,"textId":"BU_*_LOCATION_AU_bloc__ff6e04dba16ac5ef1d2e28dc22f280fc"},{"code":"Nat_Geo_Animals_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__9cdc9ed672a9d12942671255539e2f09"},{"code":"Planet_Possible_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__27d4a362f6c45d372a54b053164bfd03"},{"code":"Nat_Geo_Family_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__747bf9c4541c95d870abda6fd51c89da"},{"code":"Nat_Geo_Coronavirus_Update_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__ff6bdde8d67a24558df4344f12b185fb"},{"code":"Nat_Geo_History_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__ad40f1047ff7d20ae5cd6a84c7175eec"},{"code":"Nat_Geo_Photography_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__4a0ae93b679eac824a525b838dfbf519"},{"code":"Nat_Geo_Science_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__0d92490f28b3890ad0806108887c2275"},{"code":"Nat_Geo_Compass_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__8e8b479bd3a34136a364c83104a995b8"},{"code":"Nat_Geo_Travel_Newsletter_N","subscribed":true,"textId":"BU_*_LOCATION_*__8450564036f0849da77477398f91f6ab"}],"campaign":"Natgeo_All_Newsletters"}`, email)

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
	file, _ := os.Create("POST876")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
