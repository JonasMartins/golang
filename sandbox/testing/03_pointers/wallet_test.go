package pointers

import (
    "testing"
)

func TestWallet(t *testing.T) {

    assertBalance := func(t testing.TB, wallet Wallet, want Biticoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }

    }

    assertBalance := func(t testing.TB, wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }

    }

    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(20))
    })

    t.Run("withdraw", func(t, *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        wallet.Withdraw(10)
        assertBalance(t, wallet, Bitcoin(10))
    }

    t.Run("withdraw with no founds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err)
        assertBalance(t, wallet, startingBalance)
    })

}
