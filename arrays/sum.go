package arrays

func Sum(nums [5]int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
