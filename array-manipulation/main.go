package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	arr := make([]int, n+1)
	for i := 1; i <= m; i++ {
		var a, b, k int
		fmt.Scan(&a, &b, &k)
		for j := a; j <= b; j++ {
			arr[j] += k
		}
	}
	max := math.MinInt32
	for _, item := range arr {
		if item > max {
			max = item
		}
	}

	fmt.Println(max)
}
