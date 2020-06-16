package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IntSorted []int32

func (f IntSorted) Len() int { return len(f) }

func (f IntSorted) Less(i, j int) bool { return f[i] < f[j] }

func (f IntSorted) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	arrSorted := IntSorted(arr)
	sort.Sort(arrSorted)

	var minSum, maxSum int64 = 0, 0
	for i := 0; i < arrSorted.Len()-1; i++ {
		minSum += int64(arrSorted[i])
	}

	for i := arrSorted.Len() - 1; i > 0; i-- {
		maxSum += int64(arrSorted[i])
	}

	fmt.Printf("%d %d", minSum, maxSum)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
