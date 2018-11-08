package main

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	rotateNet := int(d) % len(a)
	// Create array the size of the initial array
	rotArr := make([]int32, len(a))
	for i, val := range a {
		if i-rotateNet < 0 {
			rotArr[i+(len(a)-rotateNet)] = val
		} else {
			rotArr[i-rotateNet] = val
		}
	}
	return rotArr

}
