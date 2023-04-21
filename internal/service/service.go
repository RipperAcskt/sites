package service

import (
	"context"
	"fmt"

	"sites/config"
	"sites/internal/model"
)

var (
	ErrEmpty = fmt.Errorf("error no rows")
)

type Service struct {
	*AuthService
	*UserService
	*Payment
	*Master
}
type Repo interface {
	AuthRepo
	UserRepo
	PaymentRepo
	MasterRepo
}
type UserRepo interface {
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetPaymentByUserId(ctx context.Context, id string) ([]model.Order, error)
}
type UserService struct {
	UserRepo
}

func New(postgres Repo, salt string, cfg *config.Config) *Service {
	return &Service{
		AuthService: NewAuthSevice(postgres, salt, cfg),
		UserService: NewUserService(postgres),
		Payment:     NewPayment(*cfg, postgres),
		Master:      NewMaster(postgres),
	}
}

func NewUserService(postgres UserRepo) *UserService {
	return &UserService{postgres}
}

func (user *UserService) GetProfile(ctx context.Context, id string) (*model.User, []model.Order, error) {
	u, err := user.GetUserById(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("get user by id: %w", err)
	}
	p, err := user.GetPaymentByUserId(ctx, id)
	if err != nil && err != ErrEmpty {
		return nil, nil, fmt.Errorf("get payment by id: %w", err)
	}
	return u, p, nil
}
