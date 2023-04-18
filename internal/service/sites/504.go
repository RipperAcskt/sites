package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST504(email string) (int, error) {
	urlString := "https://customervoice.microsoft.com/formapi/api/b66c9f75-1b5f-4d62-80ff-8ac626f15ced/users/0e41e9d4-0ef2-4d46-ae06-4b3a418a52f0/forms('dZ9stl8bYk2A_4rGJvFc7dTpQQ7yDkZNrgZLOkGKUvBUOTZPUzREUldWRkRaNlBZQUtFUU1WQjc4Mi4u')/responses"

	data := fmt.Sprintf(`{"answers":"[{\"questionId\":\"r3fb2ace1f2c6453cb8c71d310b81db73\",\"answer1\":\"%s\"},{\"questionId\":\"ree3876189c15468cb963669281b32048\",\"answer1\":\"fdsfsd\"},{\"questionId\":\"r8161dbce1310426e813a584b0f1530f7\",\"answer1\":\"fsdfsdf\"},{\"questionId\":\"rccbca8f4ca884a5d83a428d6933c2d27\",\"answer1\":\"fdsfsfdsfds\"},{\"questionId\":\"r30d7010140c34c0bb1215325f73ea0a9\",\"answer1\":\"fsdfsdfsd\"},{\"questionId\":\"rb4e83b4bf5d64421ba70fc255e7463ef\",\"answer1\":\"[\\\"Yes\\\"]\"},{\"questionId\":\"r8b96e0b791b14c6687dbb28cfeab2ba3\",\"answer1\":\"[\\\"Yes\\\"]\"},{\"questionId\":\"r781b2fc8b4f04522a82a3bcecbb5173e\",\"answer1\":\"[\\\"Access and widening participation\\\",\\\"Climate emergency\\\",\\\"Degree apprenticeships\\\",\\\"Employability and skills\\\",\\\"Equality and diversity\\\",\\\"Immigration\\\",\\\"Tackling harassment\\\",\\\"Teaching and learning\\\"]\"},{\"questionId\":\"ra8eda1c079be4f7db0ef3d35b348fa7f\",\"answer1\":null}]","startDate":"2023-03-30T16:19:52.100Z","submitDate":"2023-03-30T16:20:12.899Z"}`, email)

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
