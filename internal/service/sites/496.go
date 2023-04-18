package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST496(email string) (int, error) {
	urlString := "https://wwws.airfrance.fr/gql/v1?bookingFlow=LEISURE"

	data := fmt.Sprintf(`{"operationName":"subscribeForSalesPublicationAirline","variables":{"input":{"email":"%s","process":"BW_SUBS"},"recaptchaAssessment":{"token":"03AKH6MRHZtgtniJwtIzJxRXIkpGKi3i42iZrlZnGQTz1atsW9gbA8uIEOcxi323LsChT8e7jnhlOpp_hGUQXBQ5tLd496Pm7m0DdZyxl4zqfjjlni6J3a9iB5KZ-bcJvgclGxqOy1dUSK7BR1hNrVH-v8SE5LiPn-wFXPngEWkyUXOWPxJCod5_KO84BVkZD1Z9F0CY0eG8qV3VRFw91gwBaRwfyYJntlyef-3wGVfbUwtoLmtPvYzhEEW-XmDv_6oiCjJLTYQpsIHHXiqr_x5a5EcF_8ckjjg0KF6MB0i3rQuB34UAOY8N3mi7BL_z0cO3aFDrt8dwoZDXsrZCnAv91Cl1lXKIk4erKVATT3gVtWdk7twiQ9Akw7MHJw_tEAEUy1bfY4YtpY3_rsaMBQA9c2M1cJx5mPdfBODVklCUYB8qKUCjAzn5PTdyfoX-Z32vNc8djryF6BL9Ow5D4uv3C2S7DPKqqkbTtqn3Q0w5JxxYT2k1PO0Z85DHyDDBfY6TlSWEnGawdku0gZqTQNwK5o6tA54-Z9Dis6SeTi1lCPAbukOqX_UaK7T3Pj2vFF4MhCoxds7cOJvaU8QQlIR49crHJWmV85qAh0UDkF6LbCWBav-UwzBqj_Ea0DKc47wAkeEM2Gck9Rpw-M_1s7dbpJKUoF-bJc-5OK0thGucAGDetLx4CZZL4a9UvWpsgGKV8vxlaRoApJv8OuwR7ibxUT-pi8bjwnRvJGbcXoRgb7Te7GFp5rCLwJXyoUnjC4_Qo1cpKZpwY7yRezCjJjNH4jBi_RT5-PNJHuQz3ZuiEWvlRjZIaYcCLaB70vwEwqGGIZ3qgn_wd6ke_C4gFesRIxDjPcWzzZZPqjCd7AbrdpgXWBShxliu2I3JS0I22xdhsDt-0NcBsWxC7qK7il2nwYOwpfREbUrLTBkPm5Mm07POAyItHTx6eqqzJ9AJU_2ctIctR4M1jZ84eOOQxuSFMNaObIJw_xvHLAdhgHBmGTwazhdldMqz_PDFG1wCAhWRJViHJkUNI27Qk6Nv3xXRfG6J_VyLMYrVHRddoML5xxbQYXw07aMBfkHQ7teoOaixqhT1IYd2XS3Ooqxp5E_kDdFbfxc-JE5UTckFm2ioksXHMoFUmaKZP-veDEj-PPLebm3UMNfPbLFSvPwb3ClMK7l9I0Iq8aym1eCeRq6jABKMxz9R_F6AvIKm1XRai8Imm_Kf2oU1Gkx_UqrgGDpVFGsV7BssMgB60xFt7mwJ56UpyziyC5h2mrmfKU3Q3xtOc6cjJ13N57Bua3wfKC6J0hGVTDAImnV_fHR7MZfDPdbwRDDwVLUi4","expectedAction":"subscription"}},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"135bc2d2b8e424768ea3e1e30d5547b34dfa72b36dd5956140937636460798c6"}}}`, email)

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
