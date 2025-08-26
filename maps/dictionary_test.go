package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		compareStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		compareError(t, got, ErrNotFind)
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
func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		compareError(t, err, ErrNonExistingWord)
	})
}
func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFind {
		t.Errorf("'%s' is expected to be deleted", word)
	}
}
func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found added word")
	}
	if definition != result {
		t.Errorf("got '%s', want '%s'", result, definition)
	}
}
func compareStrings(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("got '%s', want '%s', given '%s'", result, expected, "test")
	}
}

func compareError(t *testing.T, result, expected error) {
	t.Helper()
	if result != expected {
		t.Errorf("got '%s', want '%s'", result, expected)
	}
}
