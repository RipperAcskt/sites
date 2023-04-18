package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST430(email string) (int, error) {
	urlString := "https://enterprisedunedin.outreach.co.nz/forms/subscribe"

	data := url.Values{}
	data.Set("email", email)
	data.Set("section", "all")
	data.Set("first_name", "posdafjsd")
	data.Set("last_name", "jofiasidfjosi")
	data.Set("business_name", "fjdosijfoisd")
	data.Set("location[]", "18384")
	data.Set("subscriptions[]", "115111")
	data.Set("categories[]", "57988")
	data.Set("subscriptions[]", "115175")
	data.Set("categories[]", "41627")
	data.Set("subscriptions[]", "114983")
	data.Set("categories[]", "109185")
	data.Set("topics[]", "112739")
	data.Set("topics[]", "112931")
	data.Set("topics[]", "113059")
	data.Set("topics[]", "114659")
	data.Set("topics[]", "114403")
	data.Set("topics[]", "113123")
	data.Set("topics[]", "114339")
	data.Set("topics[]", "112995")
	data.Set("topics[]", "114275")
	data.Set("terms", "on")
	data.Set("g-recaptcha-response", "03AKH6MRHoZwlMMdfa_JViRonBfz1Wvcz38WgCCKttIxdEY81f9lhftTwTmGwSaP42OML92Z35K4ROJa7gCIOz9JWGlk0vLG_62McpKpeDbClVy0EyuHB8GkSvEQMtj9I-Wh6RABWnh8XiLKaywd9rlOByNc3t5VedZ7eT6uKvkVWpbHAN4Fxe1QEs6hudPJg2vWE8bHq1kF9CU8eubMyVkl1pXUpjVEOgAwJLJ17QSD8RqNOOkj5AIdZsRagu8FofeDZw6ns60pRgRxPE_Sv6hgK6mL0twXGN9gBLTJDsQqN3t4vDMl-u4I7ktZ2DZge2telBQLnCA_EX_bS6-PwtErgRfWBGZF-_J1L_CwpGha-T7hU5o5xjxgvY3d2vUcaJfHFiQZioanx8nZTSsHvs_mWet0BXZdq8SvHP0dY9y-0gWhHXtjCP3x-JAu9koIuVZ8Rameey4Y4yFauUxg8lQdRtMq1IXPHKp9wjPlrJ66u1J0Uy6p1NuO-Y1Mq0fqFd_qhxl4cNUCE2HvNdMYshWqOgbzsLcGkWslU4bbcJbQfHFI_HdoOjm7g")

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
