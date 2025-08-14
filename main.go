package main

import (
	"fmt"
	"golang/learning/array"
)

func main() {
	// code to execute module functions
	a := []int{2, 4, 1, 6, 10, 8}
	index := array.Search1DArray(a, 10)
	fmt.Printf("Element position: %d\n", index)
}
