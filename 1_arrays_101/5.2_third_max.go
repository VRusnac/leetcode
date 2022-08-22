package arrays101

import "sort"

func thirdMax(nums []int) int {
	var count int
	in := make(map[int]bool)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i := 0; i < len(nums); i++ {
		if !in[nums[i]] {
			in[nums[i]] = true
			count++
			if count == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}
