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
	confirmError := func(t *testing.T, result error, expected error) {
		t.Helper()
		if result == nil {
			t.Fatal("expected an error but none occurred")
		}
		if result != expected {
			t.Errorf("result '%s' expected '%s'", result, expected)
		}
	}
	confirmNotError := func(t *testing.T, result error) {
		t.Helper()
		if result != nil {
			t.Fatal("unexpected error received")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(26)
		confirmBalance(t, wallet, 26)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Rodocoin(26)}
		err := wallet.Withdraw(23)
		confirmBalance(t, wallet, 3)
		confirmNotError(t, err)
	})
	t.Run("withdraw with insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Rodocoin(17)}
		err := wallet.Withdraw(25)
		confirmBalance(t, wallet, 17)
		confirmError(t, err, ErrorInsufficientBalance)
	})
}
