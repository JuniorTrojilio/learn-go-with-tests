package pointeranderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, expect Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}
	}

	confirmError := func(t *testing.T, erro error, expect string) {
		t.Helper()
		if erro == nil {
			t.Errorf("Expected an error but did not occur")
		}

		result := erro.Error()

		if result != expect {
			t.Errorf("Result '%s', expect '%s'", result, expect)
		}
	}

	t.Run("Deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Bitcoin(10))
		expect := Bitcoin(10)
		confirmBalance(t, wallet, expect)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(10))
		expect := Bitcoin(10)
		confirmBalance(t, wallet, expect)
	})

	t.Run("Withdrak without suficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		error := wallet.Withdraw(100)
		confirmBalance(t, wallet, initialBalance)
		confirmError(t, error, ErrorInsuficientBalance)
	})
}
