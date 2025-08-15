package arrays

import "testing"

func TestSum(t *testing.T) {
	nums := [...]int{1, 2, 3, 4, 5}
	result := Sum(nums)
	expected := 15
	if result != expected {
		t.Errorf("result '%d', expected '%d', given '%v'", result, expected, nums)
	}
}
