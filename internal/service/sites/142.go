package sites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func POST142(email string) (int, error) {
	urlString := "https://graphql.universal-music.de/?language=de&channel=7&channels=7"

	data := fmt.Sprintf(`{"id":"mutation0.9418551664235943","query":"mutation UnknownFile($input_0:SubscriptionByConnectArtistInput!) {subscribeByConnectArtist(input:$input_0) {clientMutationId,...F0}} fragment F0 on SubscriptionPayload {successful,error}","variables":{"input_0":{"connectLocalArtistID":620075916,"userEmail":"%s","lyticsUserId":"0c57f06c-07fc-46b1-9b2b-a9c956ffce79","referralUrl":"https://www.klassikakzente.de/newsletter/abonnieren","clientMutationId":"0"}}}`, email)

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
