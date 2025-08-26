package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyMessageIsCorrect := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Rodo", "")
		want := "Hello, Rodo!"
		verifyMessageIsCorrect(t, got, want)
	})
	t.Run("say 'Hello, world!' when a string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		verifyMessageIsCorrect(t, got, want)
	})
	t.Run("in chinese", func(t *testing.T) {
		got := Hello("Rodo", "chinese")
		want := "Nǐ hǎo, Rodo!"
		verifyMessageIsCorrect(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Rodo", "french")
		want := "Bonjour, Rodo!"
		verifyMessageIsCorrect(t, got, want)
	})

}
