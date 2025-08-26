package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, want Rodocoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}
	confirmError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("want an error but none occurred")
		}
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	confirmNotError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
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
