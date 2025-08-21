package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is just a test"
		compareStrings(t, result, expected)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, result := dictionary.Search("unknown")
		compareError(t, result, ErrNotFind)
	})
}
func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		compareError(t, err, ErrExistingWord)
		compareDefinition(t, dictionary, word, definition)
	})
}
func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found added word")
	}
	if definition != result {
		t.Errorf("result '%s', expected '%s'", result, definition)
	}
}
func compareStrings(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("result '%s', expected '%s', given '%s'", result, expected, "test")
	}
}

func compareError(t *testing.T, result, expected error) {
	t.Helper()
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
