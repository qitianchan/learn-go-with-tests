package hello;

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Elodie", "Spenish")
		want := "Helo, Elodie"

		assertCorrectMessage(t, got, want)
	})


	t.Run("saying hello to Franch", func(t *testing.T) {
		got := Hello("french", "Franch")
		want := "Bonjour, french"
		assertCorrectMessage(t, got, want)
	})
	

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	
}
