package main

import (
	"fmt"
	"sort"
)

type item struct {
	weight int
	count  int
}

var arr []item

func comparePrice(i, j int) bool {
	return float32(arr[i].weight/arr[i].count) > float32(arr[j].weight/arr[j].count)
}

func cal(W int, n int) float64 {
	sort.Slice(arr, comparePrice)
	var result float64
	for _, row := range arr {
		fmt.Println(row.weight)
		if row.count <= W {
			W -= row.count
			result += float64(row.weight)
		} else {
			result += float64(row.weight/row.count) * float64(W)
			break
		}
	}
	return result
}

func main() {
	w := 50

	items := [][]int{{100, 20}, {60, 10}, {120, 30}}
	for i := 0; i < 3; i++ {
		tmp := item{weight: items[i][0], count: items[i][1]}
		arr = append(arr, tmp)
	}
	n := len(items)
	output := cal(w, n)
	fmt.Println(output)
}
