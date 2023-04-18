package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST73(email string) (int, error) {
	urlString := "qwer"

	data := url.Values{}
	data.Set("https://sealevel.nasa.gov/newsletter_subscribe_submit", email)
	data.Set("g-recaptcha-response", "03AKH6MRHrDnl2pl_VHLingKeoMDZDDyCjWZ1j0oRlKHjS7rNhAMpwsqLoHiP3Snq1z-CvdcnmAmWt3DcN0bPNRFOA5F3ChOojugHX2guSYqDEGT7dIXJOlSrNnoOvdGB4XPnylqW6NgVxCl5lmn-iVxiv5lcjoKGsR_7NCFy4W53FqsfqL3cqo3NNVfgU4Y4jhiw0IhXTUa6HzMuTdP5IKF_V9PQnrr1PtXohIz52d7roOaWkR0A42TZsKe-QCB78SVrYm0ChspvLKQl7ntlcH_ToD8_dbrzqP4jGsI2Wa6xQdDwIribUkyhh8lvo9sx11ywvyXtL3djM0T1qI934XyzHmbY7SQCXEr-aQBtwIDRD6LjJ2wxvpMifjleMcx2qNivPyqBQl-0QuvB3bf8Yx6VymPHArl5WsdMO41qOzLisH5K1etceyigTeOM4ZwXBN9VYDy2ta0yAsyAe110M5myKeJrhMyHwkf4y5zHv7p83qbEbWutDHvtLfz9CxQpndhOlAont-bTd1Ru1RK8TWc9dRckQrNTRmQ")
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
