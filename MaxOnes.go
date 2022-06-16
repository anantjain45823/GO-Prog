package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func cal(nums []int) int {
	res := 0
	curr_max := 0
	n := len(nums)
	for j := 0; j < n; j++ {
		if nums[j] == 1 {
			curr_max++
		} else {
			res = Max(res, curr_max)
			curr_max = 0
		}
	}
	return res
}
func main() {
	items := []int{1, 1, 1, 3, 11, 1, 1, 4, 5, 2, 1}
	res := cal(items)
	fmt.Println(" MaxOnes: ", res)
}
