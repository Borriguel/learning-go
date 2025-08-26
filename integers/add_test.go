package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(7, 2)
	want := 9
	if sum != want {
		t.Errorf("want '%d', got '%d'", want, sum)
	}
}
