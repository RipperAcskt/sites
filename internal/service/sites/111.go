package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST111(email string) (int, error) {
	urlString := "https://form-backend.ws.apsis.one/form"

	data := fmt.Sprintf(`{"linkId":"xcrDfPrnJWOoSC","token":"03AKH6MRGnBxtnRsoGKny8XhtUiCKQ78_B3O4MWArVT5D8TYJ7m_VYVV38sL3549H0mH03PuhyZQkMYio2HQfpivdhXJqfIyDaO4zy0KRC1_1yTKMX4SDr4bWQbh-J-X_jRP2yedWrqK3meyn-Ruf6Tv-KBzS5th1ZHSg9Og7_vd8O8cok1GNBdtX67ttVMaT1EpBtsTDES2z8_joHmyNEC19Mg7VOdJJSjy7ePvuvyV4ZtBmKQ_Q3lz06n0XXMAMilHg0Q00s6O3d-onDzHpOSXvfheRTTTYnNuMiB4dvzoz3ABStSQUY4NjijOFwqXRmIj5_D_vvvRyqRA2jo00yqGZQAQWrJPlfb361OMBY1HEfyyoMVusKQf_4xedZOX0Mk4ngvnhrD5AFd8M4enrxT9tdr202giH8WNIyqhgYySPQGOaKoyEr7dvrRQQIXbYbqYQdaE53BttBAuEVBC73N7SwppyNnv8CMGNEVQFTofXPQ330W51nybgUq6EbA5dbRJo5mbyNID0U","values":{"kcfUHsmX0-wLIIoUUtD15":"%s","BEGvFwcOdeDCCgvcv_D5P":"qwer","fs9Z4Wq3mXuAKcNq-YrDC":"qwer","eS_6unX61tZY7NxrBqTN1":["35793"]},"referrer":"https://form.apsis.one/xcrDfPrnJWOoSC","domain":"form.apsis.one","url":"https://form.apsis.one/xcrDfPrnJWOoSC","type":"submit"}`, email)

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
