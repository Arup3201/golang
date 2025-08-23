package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b = 10, 12
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
}
