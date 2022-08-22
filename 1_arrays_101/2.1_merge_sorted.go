package arrays101

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var ii, j int
	result := make([]int, m+n)

	if m != 0 && n != 0 {
		for i := 0; i < m+n; i++ {
			if j == n {
				result[i] = nums1[ii]
				ii++
			} else if nums1[ii] <= nums2[j] && nums1[ii] != 0 {
				result[i] = nums1[ii]
				ii++
			} else {
				result[i] = nums2[j]
				j++
			}
		}

		copy(nums1, result)

	} else if m == 0 {
		copy(nums1, nums2)
	}

	return nums1
}
