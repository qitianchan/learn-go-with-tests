package main

import "testing"

func TestWallet(t *testing.T) {
	
	t.Run("Deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBallence(t, wallet, want)
		
	})

	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{20}
		err := wallet.Withdraw(Bitcoin(10))

		assertBallence(t, wallet, Bitcoin(10))	
		assertNotError(t, err)

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
		wallet := Wallet{10}
		err := wallet.Withdraw(100)
		assertBallence(t, wallet, Bitcoin(10))
		assertError(t, err, InsufficientFundsError)

	})
}

var assertBallence = func(t *testing.T, w Wallet, want Bitcoin) {
	if w.Ballence() != want {
		t.Errorf("got %s, want %s", w.Ballence(), want)
	}
}

var assertError = func(t *testing.T, got error,  want error) {
	if got == nil {
		t.Errorf("want an error but didn't got one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got.Error(), want.Error())
	}
}

func assertNotError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
