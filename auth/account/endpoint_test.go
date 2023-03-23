package account

import (
	"context"
	"testing"
)

func TestMakeSignInEndpoint(t *testing.T) {
	s := NewService()
	endpoint := makeSignInEndpoint(s)
	t.Run("valid account", func(t *testing.T) {
		req := validateSignInRequest{
			Email:    "eminetto@gmail.com",
			Password: "1234567",
		}
		_, err := endpoint(context.Background(), req)
		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}
	})
	t.Run("invalid account", func(t *testing.T) {
		req := validateSignInRequest{
			Email:    "eminetto@gmail.com",
			Password: "123456",
		}
		_, err := endpoint(context.Background(), req)
		if err == nil {
			t.Errorf("expected %v received %v", ErrInvalidUser, err)
		}
	})
}
