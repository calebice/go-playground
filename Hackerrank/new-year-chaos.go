package main

import (
	"fmt"
)

// Steps:
// Start at end. check position vs value
// If not equal, go to position add to map

// Complete the minimumBribes function below.

// 2 1 5 3 4 --> 5 swap 4, swap 3, 2 swap 1
func minimumBribes(q []int32) {
	var bribes int32
	sArr := startArr(len(q))

	// Place all values in place
	for i := 0; i < len(q); i++ {
		if sArr[i] == q[i] {
			continue
		}
		if q[i] > sArr[i] {
			diff := findDiff(sArr, q, i)
			if diff > 2 {
				fmt.Println("Too chaotic")
				return
			}
			// Update temp array
			sArr = swapArr(sArr, q, i, i+int(diff))
			bribes += int32(diff)
		}

	}

	fmt.Println(bribes)
}

func findDiff(sArr, q []int32, index int) int {
	val := q[index]
	if sArr[index+1] == val {
		return 1
	}
	if sArr[index+2] == val {
		return 2
	}
	return 3
}

func swapArr(sArr, q []int32, index, end int) []int32 {
	temp := sArr[index]
	sArr[index] = q[index]
	if end == index+1 {
		sArr[end] = temp
		return sArr
	}
	sArr[end] = sArr[index+1]
	sArr[index+1] = temp
	return sArr
}

func startArr(size int) []int32 {
	var startArr []int32
	for i := 1; i <= size; i++ {
		startArr = append(startArr, int32(i))
	}

	return startArr
}
