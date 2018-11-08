package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 10, 20}
	x := -3
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })

	if i < len(arr) && arr[i] == x {
		fmt.Printf("Value found at: %v, %v\n", i, arr[i])
	} else {
		fmt.Println("Value not found")
	}

}
