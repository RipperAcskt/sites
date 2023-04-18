package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST870(email string) (int, error) {
	urlString := "https://myaccount.nytimes.com/svc/lire_ui/authorize-email"

	data := fmt.Sprintf(`{"email":"%s","abraTests":"{\"MAPS_lireCheckboxCopy_1122\":\"3_encouragement\",\"AUTH_ssoGuardrailsFlow\":\"0_Control\"}","auth_token":"H4sIAAAAAAAAA32QQXLDIAxF78I6iTEYML5BN7mCB4NomRBggNTxpLl7cRbtqt19fUlvvvRAqhSoaELn87FUFYzyMQA6IO0dhOpMa9kMcG+WcSV5tTUnwFo81Ap53puzd3nfgXtqoqCp50JKJnrCDih9viCLMWTkEkCRfuCYWLwIQbi1lGEpKWvrGZSB/LZPvyrTYLrOt+ya81FrKlPXret6Clt1VygnHa/db5TSXWMOLrwfl+zANvGilBRDgbluCRpGx3hxe9aaL2h6IJWSd1pVF8P+hB9Ym/jrMxks5NxOv2X/TzD0fH4pzcZRyIUagUdsqZUj45j1A+VccUm5oIoSa+yim2WZUJgvSg9yGAkVsv8GNth8xZ8BAAA=","form_view":"enterEmail"}`, email)

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
	file, _ := os.Create("POST870")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
