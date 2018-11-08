package main

import (
	"fmt"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	magWords := make(map[string]int)
	noteWords := make(map[string]int)
	for _, word := range magazine {
		magWords[word] = magWords[word] + 1
	}

	for _, word := range note {
		noteWords[word] = noteWords[word] + 1
	}

	for word, count := range noteWords {
		if magWords[word] < count {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}
