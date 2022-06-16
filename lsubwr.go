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

func cal(str string) int {
	hashSet := map[byte]int{}
	longestStreak := 0
	s := 0
	n := len(str)
	for j := 0; j < n; j++ {
		_, isPresent := hashSet[str[j]]
		if isPresent {
			s = Max(hashSet[str[j]]+1, s)
		}
		hashSet[str[j]] = j
		longestStreak = Max(longestStreak, j-s+1)
	}

	return longestStreak
}

func main() {
	var Str = "TakeyouForward"
	res := cal(Str)
	fmt.Println(" Largest NonRepeating Sequence Length: ", res)
}
