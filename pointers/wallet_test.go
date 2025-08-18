package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(26)
	result := wallet.Balance()
	expected := Rodocoin(26)
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
