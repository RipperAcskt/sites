package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST526(email string) (int, error) {
	urlString := "https://www.crisisgroup.org/subscribe?ajax_form=1&_wrapper_format=drupal_ajax"

	data := url.Values{}
	data.Set("email", email)
	data.Set("region_africa", "1")
	data.Set("region_asia", "1")
	data.Set("region_europe_central_asia", "1")
	data.Set("newsletters_weekly_update", "1")
	data.Set("newsletters_crisis_watch", "1")
	data.Set("newsletters_publication_french", "1")
	data.Set("terms_of_use", "1")
	data.Set("form_build_id", "form-FNvLKh1q-voZvg6iHmHhVUlDMSQdO76NiMScI4paOr4")
	data.Set("form_id", "subscribe_form")
	data.Set("_triggering_element_name", "op")
	data.Set("_triggering_element_value", "Subscribe")
	data.Set("_drupal_ajax", "1")
	data.Set("ajax_page_state[theme]", "icg_d9")
	data.Set("ajax_page_state[libraries]", "core/drupal.ajax,core/internal.jquery.form,google_analytics/google_analytics,icg_d9/footer,icg_d9/global-styling,lazy/lazy,paragraphs/drupal.paragraphs.unpublished,stripe/stripe,system/base")

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
