package search

func LinearIterative(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func LinearRecursive(arr []string, target string, index int) int {
	if index >= len(arr) {
		return -1
	}
	if arr[index] == target {
		return index
	}
	return LinearRecursive(arr, target, index+1)
}
