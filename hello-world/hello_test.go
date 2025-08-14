package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Rodo")
	expected := "Hello, Rodo!"
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
