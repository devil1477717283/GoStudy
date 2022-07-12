package _6

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello,World"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if want != got {
			t.Errorf("got '%s',want '%s' ", got, want)
		}
	}
	assertError := func(t *testing.T, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("didn't get an error but wanted one") //如果被调用就终止测试
		}
		if err.Error() != want.Error() {
			t.Errorf("got %s want %s", err, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, 0)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(0))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}
