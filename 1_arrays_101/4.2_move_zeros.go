package arrays101

func moveZeroes(nums []int) {
	var index int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for j := index; j < len(nums); j++ {
		nums[j] = 0
	}
}
