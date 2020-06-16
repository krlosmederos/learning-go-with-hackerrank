package main

import (
	"fmt"
	"math"
)

func main() {

	arr := make([][]int, 6)
	for i := 0; i < 6; i++ {
		arr[i] = make([]int, 6)
		for j := 0; j < 6; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	result := math.MinInt32

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			sum := arr[r][c] + arr[r][c+1] + arr[r][c+2]
			sum += arr[r+1][c+1]
			sum += arr[r+2][c] + arr[r+2][c+1] + arr[r+2][c+2]

			if sum > result {
				result = sum
			}

		}
	}

	fmt.Println(result)
}
