package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5, 6}

func main() {

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	fmt.Printf("ARR: %#v", arr)
}
