package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)

	dict := make(map[string]int)
	for i := 0; i < n; i++ {
		var word string
		fmt.Scan(&word)
		_, exist := dict[word]
		if exist {
			dict[word] = dict[word] + 1
		} else {
			dict[word] = 1
		}
	}

	fmt.Scan(&q)
	results := make([]int, q)
	for i := 0; i < q; i++ {
		var word string
		fmt.Scan(&word)
		results[i] = dict[word]
	}

	for i := 0; i < q; i++ {
		fmt.Println(results[i])
	}
}
