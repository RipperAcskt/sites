package sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func POST890(email string) (int, error) {
	urlString := "https://link.dealsplus.com/graphql"

	data := fmt.Sprintf(`{"operationName":"emailSignup","variables":{"name":null,"email":"%s","password":"qwerrewq"},"query":"mutation emailSignup($name: String, $email: String!, $password: String!) {\n  signUp(name: $name, email: $email, password: $password) {\n    token\n    error\n    user {\n      id\n      gender\n      avatarType\n      name\n      emailAddress\n      level\n      __typename\n    }\n    __typename\n  }\n}\n"}`, email)

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
	file, _ := os.Create("POST890")
	body, _ := io.ReadAll(response.Body)
	file.WriteString(string(body))
	return response.StatusCode, nil
}
