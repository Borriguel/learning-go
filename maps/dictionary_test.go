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
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	expected := "this is just a test"
	result, err := dictionary.Search("test")
	compareError(t, err, nil)
	compareStrings(t, expected, result)
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
