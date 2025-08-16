package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	val := 1
	fmt.Printf("val=%d\n", val)

	incr(&val)
	fmt.Printf("val=%d\n", val)
}
