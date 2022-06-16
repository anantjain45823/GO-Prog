package main

import (
	"fmt"
)

func cal(W int, arr []int) []int {
	n := len(arr)
	var result []int
	for i := n - 1; i > 0; i-- {
		for arr[i] <= W {
			W -= arr[i]
			result = append(result, arr[i])
		}
		if W == 0 {
			break
		}
	}
	return result
}

func main() {
	var W int
	fmt.Scanf("%d", &W)

	items := []int{1, 2, 5, 10, 20, 50, 100, 500, 1000}
	output := cal(W, items)
	fmt.Println(output)
}
