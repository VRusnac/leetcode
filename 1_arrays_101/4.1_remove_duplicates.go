package arrays101

func removeDuplicates(nums []int) int {
	var compare = nums[0]
	var count int

	for i := 0; i < len(nums); i++ {
		if nums[i] != compare {
			count++
			nums[count] = nums[i]
			compare = nums[i]
		}
	}

	return count + 1
}
