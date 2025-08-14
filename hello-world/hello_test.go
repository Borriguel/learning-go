package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyMessageIsCorrect := func(t testing.TB, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Rodo", "")
		expected := "Hello, Rodo!"
		verifyMessageIsCorrect(t, result, expected)
	})
	t.Run("say 'Hello, world!' when a string is empty", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world!"
		verifyMessageIsCorrect(t, result, expected)
	})
	t.Run("in chinese", func(t *testing.T) {
		result := Hello("Rodo", "chinese")
		expected := "Nǐ hǎo, Rodo!"
		verifyMessageIsCorrect(t, result, expected)
	})
	t.Run("in french", func(t *testing.T) {
		result := Hello("Rodo", "french")
		expected := "Bonjour, Rodo!"
		verifyMessageIsCorrect(t, result, expected)
	})

}
