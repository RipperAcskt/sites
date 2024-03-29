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
	GetAllPayments(ctx context.Context) ([]model.Order, error)
	GetPaymentById(ctx context.Context, id string) (string, float64, string, error)
	SetBalance(ctx context.Context, userId string, balance float64) error
}

type Payment struct {
	PaymentRepo
	client coinbase.APIClient
	cfg    config.Config
}

func NewPayment(cfg config.Config, repo PaymentRepo) *Payment {
	return &Payment{repo, coinbase.Client(cfg.COINBASE_API), cfg}
}

func (p *Payment) CreateCharge(ctx context.Context, chargeParams coinbase.ChargeParam, id uint64) (*model.Order, error) {
	charge, err := p.client.Charge.Create(chargeParams)
	if err != nil {
		return nil, fmt.Errorf("create failed: %w", err)
	}

	data, _ := json.Marshal(charge)
	fmt.Println(string(data))
	var order model.Order
	json.Unmarshal(data, &order)
	order.State = model.Submitted
	order.UserId = id
	err = p.CreatePayment(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("create payment failed: %w", err)
	}
	return &order, nil
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

func (p *Payment) GetPayments(ctx context.Context) ([]model.Order, error) {
	orders, err := p.GetAllPayments(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all payments failed: %w", err)
	}
	return orders, nil
}

func (p *Payment) Complete(ctx context.Context, id, userId string) error {
	orderId, balance, state, err := p.GetPaymentById(ctx, id)
	if err != nil {
		return fmt.Errorf("get payment id failed: %w", err)
	}

	if state != string(model.Submitted) {
		return ErrComplete
	}

	err = p.CompletePayment(ctx, orderId, id, balance)
	if err != nil {
		return fmt.Errorf("complete payment failed: %w", err)
	}

	err = p.SetBalance(ctx, userId, balance)
	if err != nil {
		return fmt.Errorf("set balance failed: %w", err)
	}

	return nil
}
