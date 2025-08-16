package main

import "fmt"

func incr(v *int) {
	*v = *v + 1
}

func main() {
	val := 1
	fmt.Printf("val=%d\n", val)

	incr(&val)
	fmt.Printf("val=%d\n", val)
}
