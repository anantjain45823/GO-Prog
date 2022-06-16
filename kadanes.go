package main

import (
	"fmt"
)

func cal(arr []int) int {
	msf := arr[0]
	meh := 0
	s := 0
	s = s + 0
	for i := 0; i < len(arr); i++ {

		meh += arr[i]
		if meh > msf {

			msf = meh
		}
		if meh < 0 {
			meh = 0
			s = i + 1

		}
	}

	return msf
}

func main() {
	items := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	msf := cal(items)
	fmt.Println(" Sum: ", msf)
}
