package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"sites/config"
	"sites/internal/model"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrUserDoesNotExists = fmt.Errorf("user does not exists")
	ErrIncorrectPassword = fmt.Errorf("incorrect password")
)

type AuthRepo interface {
	CreateUser(ctx context.Context, user model.User) error
	CheckUserByPhoneNumber(ctx context.Context, email string) (*model.User, error)
}

type TokenRepo interface {
	AddToken(token string, expired time.Duration) error
	GetToken(token string) bool
}
type AuthService struct {
	AuthRepo
	salt string
	cfg  *config.Config
}

func NewAuthSevice(postgres AuthRepo, salt string, cfg *config.Config) *AuthService {
	return &AuthService{postgres, salt, cfg}
}

func (s *AuthService) SingUp(ctx context.Context, user model.User) error {
	var err error
	user.Password, err = s.GenerateHash(user.Password)
	if err != nil {
		return fmt.Errorf("generate hash failed: %w", err)
	}

	err = s.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GenerateHash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", fmt.Errorf("write failed: %w", err)
	}
	return string(hash.Sum([]byte(s.salt))), nil
}

func (s *AuthService) SingIn(ctx context.Context, user model.User) (*Token, error) {
	userDB, err := s.CheckUserByPhoneNumber(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("check user by phone number failed: %w", err)
	}
	if user.Email == "admin" {
		if userDB.Password == user.Password {
			params := TokenParams{
				ID:                userDB.ID,
				Type:              User,
				Admin:             true,
				HS256_SECRET:      s.cfg.HS256_SECRET,
				ACCESS_TOKEN_EXP:  s.cfg.ACCESS_TOKEN_EXP,
				REFRESH_TOKEN_EXP: s.cfg.REFRESH_TOKEN_EXP,
			}

			token, err := NewToken(params)
			if err != nil {
				return nil, fmt.Errorf("new token failed: %w", err)
			}

			return token, nil
		}
	}
	hash := sha1.New()
	_, err = hash.Write([]byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("write failed: %w", err)
	}

	if userDB.Password != string(hash.Sum([]byte(s.salt))) {
		return nil, ErrIncorrectPassword
	}

	params := TokenParams{
		ID:                userDB.ID,
		Type:              User,
		Admin:             false,
		HS256_SECRET:      s.cfg.HS256_SECRET,
		ACCESS_TOKEN_EXP:  s.cfg.ACCESS_TOKEN_EXP,
		REFRESH_TOKEN_EXP: s.cfg.REFRESH_TOKEN_EXP,
	}

	token, err := NewToken(params)
	if err != nil {
		return nil, fmt.Errorf("new token failed: %w", err)
	}

	return token, nil
}
