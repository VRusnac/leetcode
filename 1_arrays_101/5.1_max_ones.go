package arrays101

func findMaxConsecutiveOnes2(nums []int) int {
	var max int

	for i := 0; i < len(nums); i++ {
		zeros := 0

		for j := i; j < len(nums); j++ {
			if zeros == 2 {
				break
			}

			if nums[j] == 0 {
				zeros++
			}

			if zeros <= 1 {
				if max < j-i+1 {
					max = j - i + 1
				}
			}
		}
	}

	return max
}
