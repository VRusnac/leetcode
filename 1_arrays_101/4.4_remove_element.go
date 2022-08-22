package arrays101

func removeElement2(nums []int, val int) int {
	var k int

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
