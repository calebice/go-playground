package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sherlockAndAnagrams("cdcd"))
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	var total int32
	wordTable := make(map[string]int)
	for i := 0; i < len(s)/2; i++ {
		for j := i; j < len(s); j++ {
			word := s[i:j]
			// Sort the substring to verify the letters line up
			wordTable[sortString(word)] = wordTable[word] + 1
		}
	}

	for _, count := range wordTable {
		total += int32(count - 1)
	}

	return total
}

type sortRunes struct {
	runes []rune
}

func (s sortRunes) Len() int {
	return len(s.runes)
}

func (s sortRunes) Less(i, j int) bool {
	return s.runes[i] < s.runes[j]
}

func (s sortRunes) Swap(i, j int) {
	temp := s.runes[i]
	s.runes[i] = s.runes[j]
	s.runes[j] = temp
}

func sortString(s string) string {
	runes := []rune(s)
	sRunes := sortRunes{runes}
	sort.Sort(sRunes)
	fmt.Println(string(sRunes.runes))
	return string(sRunes.runes)
}
