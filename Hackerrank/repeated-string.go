package main

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	remainder := int(n) % len(s)

	count := int(n) / len(s)

	aCount := int64(count) * calculateA(s)

	remainA := calculateA(s[0:remainder])
	return aCount + remainA

}

func calculateA(s string) int64 {
	var count int64
	for _, a := range s {
		if string(a) == "a" {
			count++
		}
	}
	return count
}
