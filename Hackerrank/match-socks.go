package main

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	sockMap := make(map[int32]int)

	for _, sock := range ar {
		sockMap[sock] = sockMap[sock] + 1
	}

	var totalPairs int32
	for _, socks := range sockMap {
		totalPairs += int32(socks / 2)
	}

	return totalPairs
}
