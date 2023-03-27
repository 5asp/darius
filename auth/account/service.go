package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aszeta/darius/auth/security"
)

type Service interface {
	SignUp(ctx context.Context, account Account) int
	SignIn(ctx context.Context, mail, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("invalid account")
	ErrInvalidToken = errors.New("invalid token")
)

type service struct {
	database *sql.DB
}

func NewService() *service {
	return &service{}
}

func (s *service) SignIn(ctx context.Context, email, password string) (string, error) {
	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return "", ErrInvalidUser
	}
	config := NewConfig(".")
	token, err := security.NewToken(email, config.App.Salt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) SignUp(ctx context.Context, account Account) (result int) {
	config := NewConfig(".")
	db := NewDB(config)
	//@TODO create validation rules, using databases or something else
	account.PasswordHash, _ = security.Password(account.PasswordHash)
	// dynamic
	insertDynStmt := `insert into "account"("email", "password_hash","created_at") values($1, $2,$3)`
	_, e := db.Exec(insertDynStmt, account.Email, account.PasswordHash, account.CreatedAt)
	if e != nil {
		result = 0
		return
	}
	result = 1
	return result
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
