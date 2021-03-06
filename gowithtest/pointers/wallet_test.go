package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, want Bitcoint, wallet Wallet) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoint(10))

		assertBalance(t, Bitcoint(10), wallet)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(10)}

		err := wallet.Withdraw(Bitcoint(4))

		assertBalance(t, Bitcoint(6), wallet)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(10)}

		err := wallet.Withdraw(Bitcoint(100))
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Got un expected error: %v", err)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Error("Should return error")
	}
	if err != want {
		t.Errorf("want error: %v, but got: %v", want, err)
	}
}
