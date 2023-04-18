package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST187(email string) (int, error) {
	urlString := "https://forms.hscollectedforms.net/collected-forms/submit/form"

	data := fmt.Sprintf(`{"contactFields":{"email":"%s"},"formSelectorClasses":"","formSelectorId":"#frm-newsletterForm","formValues":{"Мною прочитана Политика конфиденциальности Ataccama.":"Checked","Я хочу получать новости, приглашения на будущие мероприятия и другие предложения от Ataccama.":"Not Checked","Я даю согласие, чтобы партнеры Ataccama, выступающие в качестве контролеров данных, отправляли мне новости, приглашения на будущие мероприятия и другие предложения.":"Checked"},"labelToNameMap":{"Мною прочитана Политика конфиденциальности Ataccama.":"privacy_policy","Я хочу получать новости, приглашения на будущие мероприятия и другие предложения от Ataccama.":"subscribe","Я даю согласие, чтобы партнеры Ataccama, выступающие в качестве контролеров данных, отправляли мне новости, приглашения на будущие мероприятия и другие предложения.":"contact_agreement"},"pageTitle":"Подписаться на рассылку | Ataccama","pageUrl":"https://www.ataccama.com/ru/contact/sign-up-for-newsletter","portalId":23390279,"type":"SCRAPED","utk":"0e6e5e5a1ac69baa931f5e7585cf73ab","uuid":"54b16266-5774-417b-a53f-114b1dfca60d","version":"collected-forms-embed-js-static-1.331","collectedFormId":"frm-newsletterForm","collectedFormAction":"/ru/contact/sign-up-for-newsletter"}`, email)

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
