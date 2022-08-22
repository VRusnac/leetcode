package arrays101

func findNumbers(nums []int) int {
	var total int

	for i := 0; i < len(nums); i++ {
		var tmp int

		for nums[i] > 0 {
			nums[i] = nums[i] / 10
			tmp++
		}

		if tmp%2 == 0 {
			total++
		}
	}
	return total
}
