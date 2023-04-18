package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sites/config"
	"sites/internal/model"
	"strconv"

	"github.com/02amanag/coinbase-commerce-go"
)

type PaymentRepo interface {
	CreatePayment(ctx context.Context, order model.Order) error
	GetPaymentId(ctx context.Context, id string) (string, float64, error)
	CompletePayment(ctx context.Context, order_id, user_id string, balance float64) error
	SetState(ctx context.Context, order_id string) error
}

type Payment struct {
	PaymentRepo
	client coinbase.APIClient
	cfg    config.Config
}

func NewPayment(cfg config.Config, repo PaymentRepo) *Payment {
	return &Payment{repo, coinbase.Client(cfg.COINBASE_API), cfg}
}

func (p *Payment) CreateCharge(ctx context.Context, chargeParams coinbase.ChargeParam, id uint64) error {
	charge, err := p.client.Charge.Create(chargeParams)
	if err != nil {
		return fmt.Errorf("create failed: %w", err)
	}

	data, _ := json.Marshal(charge)

	var order model.Order
	json.Unmarshal(data, &order)
	order.State = model.Submitted
	order.UserId = id
	err = p.CreatePayment(ctx, order)
	if err != nil {
		return fmt.Errorf("create payment failed: %w", err)
	}
	return nil
}

func (p *Payment) GetCharge(ctx context.Context, id string) (model.State, error) {
	orderId, balance, err := p.GetPaymentId(ctx, id)
	if err != nil {
		return model.Submitted, fmt.Errorf("get payment id failed: %w", err)
	}

	charge, err := p.client.Charge.Get(orderId)
	if err != nil {
		return model.Submitted, fmt.Errorf("create failed: %w", err)
	}

	data, _ := json.Marshal(charge)
	var order model.Order
	json.Unmarshal(data, &order)

	var sum float64
	for _, p := range order.Data.Payments {
		transactionAmount, err := strconv.ParseFloat(p, 64)
		if err == nil {
			sum += transactionAmount
		}
	}

	for _, t := range order.Data.Timeline {
		if t.Status == "EXPIRED" {
			p.SetState(ctx, orderId)
			return model.Expired, nil
		}
	}

	if sum < balance {
		return model.Submitted, nil
	}

	err = p.CompletePayment(ctx, orderId, id, balance)
	if err != nil {
		return model.Submitted, fmt.Errorf("complete payment failed: %w", err)
	}

	return model.Completed, nil
}
