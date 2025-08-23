package main

import "fmt"

func isPrime(x int) bool {
	if x==2 {
		return true
	}

	for i:=2; i*i<=x;i++ {
		if x%i==0 {
			return false
		}
	}

	return true
}

func main() {
	tests := []int{2, 3, 5, 6, 9, 11, 15, 17, 25, 27, 31, 40, 90, 111, 231, 331, 433, 470, 511, 581, 781, 989}
	for _, n := range tests {
		fmt.Printf("isPrime(%d)=%t\n", n, isPrime(n))
	}
}
