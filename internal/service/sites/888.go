package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST888(email string) (int, error) {
	urlString := "https://account.booking.com/api/identity/authenticate/v1.0/enter/email/submit?op_token=EgVvYXV0aCKpBQoUdk8xS2Jsazd4WDl0VW4yY3BaTFMSCWF1dGhvcml6ZRo1aHR0cHM6Ly9zZWN1cmUuYm9va2luZy5jb20vbG9naW4uaHRtbD9vcD1vYXV0aF9yZXR1cm4qyARVck1ENXNEQ1A3czNoLTdTQ0prMWFWZTNSSG95cTlpZXluUFllZTFMZ2NZYUJLVjRvYU83T3RISzA1eFNjaHdZdE5iaWVTWkMwZ29pX1A0SHkweDJBYV9aS1JyN2tDdy04X1NsMXg5QzdQbUx5MmVhUTd6ZmYtTGpDbXRaZlZQS0R1ZEJqTDB2X3FocGJ1a2lKV19JYU5mYzhhT3RQREJoLXdwRy1veldjOFdCV2tLUERSa3VwSzBLNTdrT241STBfbzNtUFhzV0U4ZWJaaE1HNHF6VEIwd0dOS19UbnJqTXVTekFUZkl4U0UxNDBwc0ZaZ1J4eGE0UnpfZEFvVGhqMkRIVmQtQUdKeS1LOGI1Q0Y5TkxfVERHRVQ4cDFZd1V5aXBEdVByQ25saVVOUERiLXd1ZktONTEzWXI1bnZHMGdyN0lmRkFBYVVRRE5IaTNrcDFqNHFNbm1QTUJUc3VCemxQX29ObFYzWVh5cFRLYmlscVZFMU50NURjSFJ0enhpSExPSnQ4dmQxQkktbDNQYXdGaDltNUZHU1VzWkZYcUNfQjEzeWJwV1E3UGJlZUVyaGJJeVFubzVSeWlENDEtQ0IxVFF2Slg5SGxWQ05YMzducE5oNjg3Znd3ckpuaktuaFBtMHR2MTJmUmY5SFAtTVZ3OE8wWHhlUFl0bEN5RjFPWkhwVzNvVXJEaV9tc3NTdVZzeXBFa0Zyd3FyUWs3RG05bUdpd0htLUx6eU5oRWQ3WXdJeXFLRHdnQ3FIQjNET3F1TlpLY0IEY29kZSoWCI7IEjDB5b6dwZkmOgBCAFiF2YihBg"

	data := fmt.Sprintf(`{"context":{"value":"UoEBf2GfJEAAopmn8662N2k6NUhitK73VddwrsWZGmWB_eepA0LHDSKG2rIz7U8q9W7d1KODHFyZbiLDcOCfXUHjVUMLQftPRMQBfLV851Jx7qm7iXRPtVJgzd8vNhkVAeWbsovjGACXY8kidyDoyI6t-P8iJgJkS1kSe7M4_1ZUN2yr"},"identifier":{"type":"IDENTIFIER_TYPE__EMAIL","value":"%s.com"}}`, email)

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
	file, _ := os.Create("POST888")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
