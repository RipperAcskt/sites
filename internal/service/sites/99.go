package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST99(email string) (int, error) {
	urlString := "https://warmquiver.com/v2eatA7eyQ_vqfTN9ftDIzJJlSISLxqruUqtOQm6ik7C_r-Iy8JwttfznJji26KGcWq0FMYMoVQkm"

	data := fmt.Sprintf(`{"bl":true,"abl":true,"bls":160,"flags":8715,"dt":173,"uri":"https://www.defensenews.com/air/2022/12/16/gripen-officially-joins-brazils-operational-fighter-fleet/","pv":"0u8dc70577ce8711ed8cb0b938a8fe26f3","pid":"A-5BFEC5776DDF1A0BB07337DF-7","jsv":"2.31.0","utco":10800,"ipr":true,"sid":"5-48b9c7bda7e24513bf544e107f2e7605-6763652d6575726f70652d7765737431-3","aid":"2-f2dabaa310ca2267-8dd73ad8-ce87-11ed-af35-10dbce2c83b7","email":"%s","offerID":"5cb7ce7fcfc37d6359a50568","addonOfferIDs":[],"contactAttributes":{"name":{"given":"","family":""},"company":""},"metadata":{"sourceURL":"https://www.defensenews.com/air/2022/12/16/gripen-officially-joins-brazils-operational-fighter-fleet/"}}`, email)

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
