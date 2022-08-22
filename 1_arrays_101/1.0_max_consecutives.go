package arrays101

func findMaxConsecutiveOnes(nums []int) int {
	var max, j int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			j++
			if j > max {
				max = j
			}
		} else {
			j = 0
		}
	}
	return max
}
