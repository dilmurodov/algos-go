package main

import (
	"fmt"
	"math"
)

var (
	max int = math.MinInt64
	min int = math.MaxInt64
)

var arr = []int{1, 2, 3, 4, 5, 5}

func main() {
	for _, v := range arr {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	fmt.Println(max, min)
}
