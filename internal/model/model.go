package model

import "github.com/02amanag/coinbase-commerce-go"

type State string

var (
	Submitted State = "submitted"
	Completed State = "completed"
	Expired   State = "expired"
)

type User struct {
	ID       uint64  `json:"-"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Balance  float32 `json:"balance"`
}

type Charge struct {
	Name         string
	Description  string
	Local_price  coinbase.Money
	Metadata     coinbase.Metadata
	Redirect_url string
	Cancel_url   string
}

type Order struct {
	ID     int
	UserId uint64
	State  State
	Data   struct {
		OrderId   string `json:"id"`
		HostedURL string `json:"hosted_url"`
		Pricing   struct {
			Local struct {
				Amount string `json:"amount"`
			} `json:"local"`
		} `json:"pricing"`
		Payments []string `json:"payments"`
		Timeline []struct {
			Status string `json:"status"`
		} `json:"timeline"`
	} `json:"data"`
}
