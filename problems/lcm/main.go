package main

import "fmt"

func gcd(a, b int) int {
	if b==0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a*b) / gcd(a, b)
}

func main() {
	var a, b = 3, 9
	fmt.Printf("lcm(%d, %d) = %d\n", a, b, lcm(a, b))
}
