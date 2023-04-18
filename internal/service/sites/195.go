package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST195(email string) (int, error) {
	urlString := "https://meduza.io/digests/daily/ru/captcha"

	data := url.Values{}
	data.Set("email", email)
	data.Set("token", "03AKH6MRE-4rUlpvuSm80R-0gUT4D9_vtWag1t5N0zsZUt_vInowVnGRTuRdaQKvnq-Qzv5qGBdMn92e84BxZWpXRbW3inpKuoSGxNO1iRLt7FAYWJtUjTJ3rI_mt37hCX5sJ_BMP2QBN-yWROqFZmbrHM2ZbWrY_u_aE_wNteDUL7v5wSM1vJDzy-8z90OEgI9qEVGFG4_UBYQjYFNvsyPasJXYMQPQuyASAcu_z1e0V8dnp9dOqUJMgzKHxOvDifHqtBxCl0oOo2-wVnkGNOZeU_oHL-oemdEuC5kRoyeLcKuRtmy4vkgEDbD3A_ajW6vIA7ZfoEQCwLfpeIFpkSHXpa3aKfh8H14wiQRG6_HeS-e3eC0Kq2OIHJYWqOCzbU4iM9xv-0mn-aLWVyONMZpsQKWeRyQ4_wcuA2A1YMBnIK7y2N4jrHBbBJ5zFlrxDsLByBz4hmEGHX-dXIL930cUwu_9jAHNXroTBm3F9_rvifLt8ep6ejiqh8s9LCU_7KZiPeu4n-uV5U")
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
