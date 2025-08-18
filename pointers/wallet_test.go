package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, expected Rodocoin) {
		t.Helper()
		result := wallet.Balance()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}
	confirmError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("Expected an error but none occurred")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(26)
		confirmBalance(t, wallet, 26)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Rodocoin(26)}
		wallet.Withdraw(23)
		confirmBalance(t, wallet, 3)
	})
	t.Run("withdraw with insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Rodocoin(17)}
		err := wallet.Withdraw(25)
		confirmBalance(t, wallet, 17)
		confirmError(t, err)
	})
}
