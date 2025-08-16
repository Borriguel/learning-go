package slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("colection of 5 nums", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		result := Sum(nums)
		expected := 15
		if result != expected {
			t.Errorf("result '%d', expected '%d', given '%v'", result, expected, nums)
		}
	})

	t.Run("colection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := Sum(nums)
		expected := 55
		if result != expected {
			t.Errorf("result '%d', expected '%d', given '%v'", result, expected, nums)
		}
	})
}
