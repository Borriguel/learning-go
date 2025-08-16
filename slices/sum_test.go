package slices

import (
	"reflect"
	"testing"
)

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

	t.Run("sum all collections", func(t *testing.T) {
		result := SumAll([]int{3, 9}, []int{16, 1, 4})
		expected := []int{12, 21}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})

	t.Run("sum last num of collection", func(t *testing.T) {
		result := SumLast([]int{0, 5}, []int{12, 0, -3})
		expected := []int{5, -3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})

	t.Run("sum empty slices", func(t *testing.T) {
		result := SumLast([]int{}, []int{0, -8})
		expected := []int{0, -8}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	})
}
