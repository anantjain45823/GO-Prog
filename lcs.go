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

func cal(arr []int) int {
	hashSet := map[int]bool{}
	for _, j := range arr {
		hashSet[j] = true
	}

	longestStreak := 0

	for _, j := range arr {
		if !hashSet[j-1] {
			currentNum := j
			currentStreak := 1

			for hashSet[currentNum+1] {
				currentNum += 1
				currentStreak += 1
			}

			longestStreak = Max(longestStreak, currentStreak)
		}
	}

	return longestStreak
}

func main() {
	var items = []int{3, 8, 5, 7, 6}
	res := cal(items)
	fmt.Println(" Largest Sequence Length: ", res)
}
