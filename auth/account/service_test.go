package account_test

import (
	"context"
	"testing"

	"github.com/aszeta/darius/auth/account"

	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	service := account.NewService()
	t.Run("invalid account", func(t *testing.T) {
		_, err := service.SignIn(context.Background(), "eminetto@gmail.com", "invalid")
		assert.NotNil(t, err)
		assert.Equal(t, account.ErrInvalidUser, err)
	})
	t.Run("valid account", func(t *testing.T) {
		token, err := service.SignIn(context.Background(), "eminetto@gmail.com", "1234567")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
}
