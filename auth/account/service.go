package account

import (
	"context"
	"errors"

	"github.com/aszeta/darius/auth/security"
)

type Service interface {
	SignUp(ctx context.Context, account Account) (string, error)
	SignIn(ctx context.Context, mail, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) SignIn(ctx context.Context, email, password string) (string, error) {
	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return "", ErrInvalidUser
	}
	token, err := security.NewToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) SignUp(ctx context.Context, account Account) (string, error) {
	//@TODO create validation rules, using databases or something else

	token, err := security.NewToken("1")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return "", ErrInvalidToken
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return "", ErrInvalidToken
	}
	return tData["email"].(string), nil
}