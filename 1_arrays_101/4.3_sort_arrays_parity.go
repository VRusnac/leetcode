package arrays101

func sortArrayByParity(nums []int) []int {
	var odd []int
	var rSide int

	for i := 0; i < len(nums); i++ {

		if nums[i]%2 == 0 {
			nums[rSide] = nums[i]
			rSide++
		} else {
			odd = append(odd, nums[i])
		}
	}

	for i := 0; i < len(odd); i++ {
		nums[rSide+i] = odd[i]
	}

	return nums
}
