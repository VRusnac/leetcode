package arrays101

func sortedSquares(nums []int) []int {
	var left, right = 0, len(nums) - 1
	var ans = make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		rSide := nums[right] * nums[right]
		lSide := nums[left] * nums[left]

		if lSide < rSide {
			ans[i] = rSide
			right--
		} else {
			ans[i] = lSide
			left++
		}
	}

	return ans
}
