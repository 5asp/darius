package account

import (
	"context"
	"time"

	"github.com/aszeta/darius/auth/security"
	"github.com/go-kit/kit/endpoint"
)

// 注册登录模块
type validateSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type validateSignResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeSignInEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateSignInRequest)
		token, err := svc.SignIn(ctx, req.Email, req.Password)
		if err != nil {
			return validateSignResponse{"", err.Error()}, err
		}
		return validateSignResponse{token, ""}, err
	}
}

type validateSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type validateSignUpResponse struct {
	Uuid string `json:"uuid,omitempty"`
	Err  string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeSignUpEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateSignUpRequest)
		password, err := security.Password(req.Password)
		if err != nil {
			return validateSignUpResponse{"", err.Error()}, err
		}
		data := Account{
			Email:        req.Email,
			PasswordHash: password,
			CreatedAt:    time.Now().Unix(),
		}
		result := svc.SignUp(ctx, data)
		if result == 0 {
			return validateSignUpResponse{"", err.Error()}, err
		}
		return validateSignUpResponse{data.Email, ""}, err
	}
}

// 校验令牌模块
type validateTokenRequest struct {
	Token string `json:"token"`
}

type validateTokenResponse struct {
	Email string `json:"email,omitempty"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeValidateTokenEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateTokenRequest)
		email, err := svc.ValidateToken(ctx, req.Token)
		if err != nil {
			return validateTokenResponse{"", err.Error()}, err
		}
		return validateTokenResponse{email, ""}, err
	}
}
