package binarysearch

// BinarySearch uses binary search to check for presence of a value in a sorted array
func BinarySearch(arr []int, val int) bool {
	return binarySearch(arr, val, len(arr)/2, len(arr)/2)
}

// Terminating conditions:
// if

// [0, 1, 5, 6, 10]
// Start at 5, looking for 10
// Index = 2: prev = 0
// Index = 2 + (5 - 2)/2 = 2 + 1 = 3 prev = 2
// Index = 3 + (5-3)/2 = 4 prev = 3
// Index = 4 Found!
// Start at 5, looking for 12
// Index = 2: prev = 0
// Index = 2 + (5 - 2)/2 = 2 + 1 = 3 prev = 2
// Index = 3 + (5-3)/2 = 4 prev = 3
// Index = 4 + (5-4)/2 = 5? return false if index > len-1

// [0, 2, 5, 6, 9] looking for 3
// return false if:
// Index is > len-1
// arr[prev] < val < arr[index]
// arr[index] > val > arr[prev]

func binarySearch(arr []int, val, index, prev int) bool {
	if arr[index] == val {
		return true
	}
	if arr[index] > val {
		// Check if the curr index is larger than val
		if arr[prev] < val {
			return false
		}
		index = index / 2
		prev = index
		return binarySearch(arr, val, index, prev)
	}
	prev = index
	index = index + (len(arr)-index)/2
	return binarySearch(arr, val, index, prev)

}
