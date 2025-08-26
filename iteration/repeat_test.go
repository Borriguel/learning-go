package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a", 6)
	want := "aaaaaa"
	if repetitions != want {
		t.Errorf("want '%s', got '%s'", want, repetitions)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
