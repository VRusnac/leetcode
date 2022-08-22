package arrays101

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j]*2 || arr[i]*2 == arr[j] {
				return true
			}
		}
	}

	return false
}
