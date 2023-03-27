package account

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	a := NewConfig("..")
	t.Log("app:", a.App.Name)
	t.Log("port:", a.App.Addr)
}
