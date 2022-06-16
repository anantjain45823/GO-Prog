package main

import (
	"fmt"
	"sort"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func countPlatforms(n int, arr []int, dep []int) int {
	sort.Ints(arr)
	sort.Ints(dep)

	ans := 1
	count := 1
	i := 1
	j := 0
	for i < n && j < n {
		if arr[i] <= dep[j] {
			count++
			i++
		} else //one platform can be reduced
		{
			count--
			j++
		}
		ans = max(ans, count)
	}
	return ans
}

func main() {
	arr := []int{900, 945, 955, 1100, 1500, 1800}
	dep := []int{920, 1200, 1130, 1150, 1900, 2000}
	n := len(dep)
	res := countPlatforms(n, arr, dep)
	fmt.Println(res)

}
