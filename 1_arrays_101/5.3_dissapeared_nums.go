package arrays101

import "sort"

func findDisappearedNumbers(nums []int) []int {
	var tmp []int
	sort.Ints(nums)
	in := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		in[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			if !in[i+1] {
				tmp = append(tmp, i+1)
			}
		}
	}
	return tmp
}
