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
}
