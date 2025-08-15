package main

import (
	"fmt"

	"example.com/add"
	"example.com/multiply"
)

func main() {
	fmt.Println(add.AddInts(2, 3) - multiply.MultiplyInts(4, 2))
}
