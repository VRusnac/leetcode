package arrays101

import (
	"sort"
)

func heightChecker(heights []int) int {
	var total int

	expected := make([]int, len(heights))
	copy(expected, heights)

	sort.Ints(expected)

	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			total++
		}
	}

	return total
}
