package basic

import "fmt"

func PrintHelloWorld() {
	fmt.Println("Hello World")
}

func Swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
