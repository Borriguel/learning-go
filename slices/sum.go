package slices

func Sum(nums []int) (sum int) {
	for x := range nums {
		sum += nums[x]
	}
	return
}
