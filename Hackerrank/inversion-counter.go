package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countInversions function below.
func countInversions(arr []int32) int64 {
	return inversionCount(arr, 0, len(arr), 0)
}

func inversionCount(arr []int32, leftStart, rightEnd int, invCount int64) int64 {
	middle := (rightEnd - leftStart) / 2

	if leftStart < rightEnd {
		invCount += inversionCount(arr, leftStart, middle, invCount)
		invCount += inversionCount(arr, middle+1, rightEnd, invCount)
	}

	count := handleSort(arr, leftStart, rightEnd, middle)

	return count
}

func handleSort(arr []int32, leftStart, rightEnd, middle int) int64 {
	var count int64
	for i := leftStart; i < rightEnd; i++ {
		for j := i + 1; j < rightEnd; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = arr[i]
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := countInversions(arr)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
