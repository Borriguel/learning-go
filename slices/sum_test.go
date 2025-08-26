package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("colection of 5 nums", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15
		if got != want {
			t.Errorf("got '%d', want '%d', given '%v'", got, want, nums)
		}
	})

	t.Run("colection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Sum(nums)
		want := 55
		if got != want {
			t.Errorf("got '%d', want '%d', given '%v'", got, want, nums)
		}
	})

	t.Run("sum all collections", func(t *testing.T) {
		got := SumAll([]int{3, 9}, []int{16, 1, 4})
		want := []int{12, 21}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%d', want '%d'", got, want)
		}
	})

	t.Run("sum last num of collection", func(t *testing.T) {
		got := SumLast([]int{0, 5}, []int{12, 0, -3})
		want := []int{5, -3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%d', want '%d'", got, want)
		}
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumLast([]int{}, []int{0, -8})
		want := []int{0, -8}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%d', want '%d'", got, want)
		}
	})
}
