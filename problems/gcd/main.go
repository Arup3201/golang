package main

import "fmt"

func gcd(a, b int) int {
	if b==0 {
		return a
	}

	return gcd(b, a%b)
}

func main() {
	var a, b = 4, 10
	fmt.Printf("gdc(%d, %d) = %d\n", a, b, gcd(a, b))
}
