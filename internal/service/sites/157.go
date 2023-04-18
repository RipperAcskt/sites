package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST157(email string) (int, error) {
	urlString := "https://www.vzbv.de/politik/themen-newsletter?ajax_form=1&_wrapper_format=drupal_ajax"

	data := url.Values{}
	data.Set("email", email)
	data.Set("ajax_page_state[libraries]", "classy/base,classy/messages,core/internal.jquery.form,core/normalize,fontawesome/fontawesome.svg,fontawesome/fontawesome.svg.shim,formtips/formtips,image_sizes/core,paragraphs/drupal.paragraphs.unpublished,shariff/shariff,shariff/shariff-naked,social_media_links/social_media_links.theme,statistics/drupal.statistics,system/base,vzbv/global-styling")
	data.Set("_triggering_element_name", "op")
	data.Set("_triggering_element_value", "Anmelden")
	data.Set("_drupal_ajax", "1")
	data.Set("ajax_page_state[theme]", "vzbv")
	data.Set("form_build_id", "form-78w0tkltj8yhbxQU4Y9MSNGjAH4mgVdsWWlbdhK-pc8")
	data.Set("form_id", "full_subscription_form")
	data.Set("lebensmittel", "1")
	data.Set("ePresseschau", "1")
	data.Set("verbraucherschutz-aktiv", "1")
	data.Set("salut", "female")
	data.Set("firstname", "qwer")
	data.Set("lastname", "qwer")
	data.Set("ajax_form", "1")
	data.Set("_wrapper_format", "drupal_ajax")
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
