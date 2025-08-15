package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(7, 2)
	expected := 9
	if sum != expected {
		t.Errorf("expected '%d', result '%d'", expected, sum)
	}
}
