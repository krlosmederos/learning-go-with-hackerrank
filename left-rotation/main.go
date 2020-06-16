package main

import (
	"fmt"
)

// func solve(arr []int32) []int32 {

// }

func main() {
	var len, d int

	fmt.Scan(&len, &d)
	arr := make([]int, len)

	for i := 0; i < len; i++ {
		fmt.Scan(&arr[i])
	}

	rounds := d % len
	arr1 := arr[0:rounds]
	arr2 := arr[rounds:len]
	result := append(arr2, arr1...)
	for i := 0; i < len; i++ {
		if i != len-1 {
			fmt.Printf("%d ", result[i])
		} else {
			fmt.Printf("%d", result[i])
		}
	}
	fmt.Println()
}
