package arrays101

func GetMax(arr []int) int {
	var max int

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func replaceElements(arr []int) []int {
	var length = len(arr)

	for i := 0; i < length-1; i++ {
		arr[i] = GetMax(arr[i+1 : length])
	}

	arr[length-1] = -1

	return arr
}
