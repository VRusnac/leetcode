package arrays101

func validMountainArray(arr []int) bool {
	var increasing = true
	var top int

	for i := 0; i < len(arr)-1; i++ {

		if arr[i] == arr[i+1] {
			return false
		}

		if increasing {
			if arr[i] > arr[i+1] {
				increasing = false
				top = i
			}
		} else {
			if arr[i] < arr[i+1] {
				return false
			}
		}
	}

	if top == 0 || top == len(arr)-1 {
		return false
	}

	return true
}
