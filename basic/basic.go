package basic

import "fmt"

func PrintHelloWorld() {
	fmt.Println("Hello World")
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func AddIntegers(a int, b int) int {
	return a + b
}
