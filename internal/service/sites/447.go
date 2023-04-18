package sites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func POST447(email string) (int, error) {
	urlString := "http://www.revistagastroenterologiamexico.org/en-alerta-email"

	data := url.Values{}
	data.Set("txtEmail", email)
	data.Set("txtNombre", "fsadfas")
	data.Set("txtApellido1", "fasdfasdfas")
	data.Set("txtApellido2", "fasdfasd")
	data.Set("estadoProfesional", "Profesional")
	data.Set("cmbPais", "ALBANIA")
	data.Set("info_elsevier", "acepted")
	data.Set("info_partners", "acepted")
	data.Set("clausula", "acepted")
	data.Set("g-recaptcha-response", "03AKH6MRHbq201k9uWcfnHq-xPqA0o9xK0N0p9AYiO0LsR_SZRF-ZIjtiTDH6N-2s_FdvyHYdunF8k9Aq9PRljRbhvtstNaVMkUXCO3elv9wGfmlmSGTIwZlIF0Kwpar4LrNDD-OMtgoAwLM8wOOTcbFePfdyDD1h-ID8fwtukNPdf3OzO6ypG1Q8GSlJHfrp4SYvpazhxbRDW1LOK9JJ1hmaDZWS-zFCrsr7l2Uirl8H67ZTIGnOJxT1HnXV8VVSCDElydQV861siTITLbFYHo8x_F4daVeNMsdEXiFfRHgoMeFBbcfLDNu-xffHIvGGLcMjsgPRsFViVeY2CzNK3hbeXKQ9dyM-XA2HYu05MDpyvu9ltdwXCQ2Gw4ewfqM7_yf6fzvQbPzBGoSaZDPGbk6yMs3arCeR5ybTaCb4QabBZXJ0MZnRmH8O8EQ3BnFaFQZIpUMwg4d5n_meXPITA_WbDQV2Oy5mRr9f0qBGmcnOQN54VFEtgvLsPForJ17s7A6at-948fqUrhcWZ50kOEz06dwqT02XlBjR4lKpRDpzkJS3xwmbIcBJ6fuLCh2Eg1fBZZozBzBlj")

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
