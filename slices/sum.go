package slices

func Sum(nums []int) (sum int) {
	for x := range nums {
		sum += nums[x]
	}
	return
}

func SumAll(nums ...[]int) (sum []int) {
	for _, num := range nums {
		sum = append(sum, Sum(num))
	}
	return
}

func SumLast(nums ...[]int) (sum []int) {
	for _, num := range nums {
		if len(num) == 0 {
			sum = append(sum, 0)
		} else {
			final := num[1:]
			sum = append(sum, Sum(final))
		}
	}
	return
}
