package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST191(email string) (int, error) {
	urlString := "https://api.y-r.by/api/v1/subscriptions"

	data := fmt.Sprintf(`{"last_name":"qwer","appeal":"Mrs","name":"qwer","is_receive_more_interesting_offer":true,"email":"%s","token":"03AKH6MRF1hmyroEr0w-gaJ0v3dtz8CJYPeIpGA-ydYHSlTY_TW3fzfuqAbfVqaGky2ASEXKoxybMJGk2aV-FWNbhQMGU7xkYYwYqo6UH0xsrakAp_wFCiVZsUqw3ew4Lm9fLfiqEqYlvT8ZgMKCF8wKIxTEdxROMs_frsE1R8HYfGRE-_XcvzfVIzNiq6pIzS5tddB_3TmK87q98KaAVp6cRne8TVtVSGmNaJgySXguPji_wQxNc23c5w7EE3GqWycZoUz9afJLM9-RhgbhC0Fcf9YmLDqXSe1xIrvKnd560W5YLCH5m6OuhNS3WShcP5BZ_a8JdJLuW74ucMCPIQBx4pWDZQByzd9SeSR3eRTfGv8t-dpb7TLQT4KXjY64spwmU16WxuAsIqMFwJjqGek06O8mUdHd1mOTyv-ovJyfdnYCGkb33EdOaAc5a2bgcSf4b2eDbh3n9o78VkdlWm6RQSVbCWUN6-jGDTmCPOPHSpX1tPLA7uJTM5bt0EHMgIi-_M47YsTmEeRpcYxuJ3X-O0Qv8NPtuaH5WNXwbOCT_eeoNy29gn6gEUelUifc5LkJdRp3aIp5k_p5fBnUCzamRVa0s7q6ilej3MeeZqbcFy3fmLK9JTX3c20rzXjUUH1EDOVWjYazHAGt9i2NUoAuJxwMCr_Wx8H_eHC99HsYjpeTx4hTaxlQ2EgNt2IGbHhXM6wQTWcuN3cJ1BemymwUMxOtZo-WZyjzDz5P_Ycj8wxjS0Br1a7xZgikAPLole8Pn2c3OCOTtFDOM1Xz_vN3nV6Agn2rWTI5_rw0WJVBTP-xC2-XqzDIrWgX2CSNb2tZ-X0xDyMxIL4Exi19wcpUDM5vlLZaiDN67gNuLVYyiXA1uoEiEEeEhjdoRSD9GPaXVWTzDLPmATxS4S77BlTsTqoddo2uNbUSWbOhzV9JqHNqjDSpltev1FNMJQSou7NwGsBtCDJ8d2UYlThu_rt-QKJNd3gUEWQRKJC6rLtt_szdgYemFuanjMKUSjeyitMfPOfDx4ToABkjGBOB97YP1Fo_IJM0JafppUD2cGh7r6DvruggkjJChn5cF23jf1wBTtRn1WlAIPgZvg_GlXwsr1LnNjLAlziE4AJLwUnXaIbcORY65E8ef139D3Ebw80c9MIuvkGudXJr6Zi4g-Lh8XaubvZlLUFK7NK0Asij6X5HEBZEc78up5OFBdaj2yD1OyxtMoiEbptppA-JfC0GEsozI7cIThLlcXu6EpwYnr52B9gjGcFJXo5X27Jczzh17QfypdmxS41XM45V3XIqQc0ztOXsD4JIc8mSipNaYWFQSEhwiGfOQ"}`, email)

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
