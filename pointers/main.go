package main

import "fmt"

func zeroval(iVal int) {
	iVal = 0
}

func zeorptr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	zeroval(i)
	fmt.Println("after calling zeroval i =", i)

	zeorptr(&i)
	fmt.Println("after calling zeorptr i =", i)
}
