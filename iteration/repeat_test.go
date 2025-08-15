package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a")
	expected := "aaaaa"
	if repetitions != expected {
		t.Errorf("Expected '%s', result '%s'", expected, repetitions)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
