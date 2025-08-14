package basic

import (
	"fmt"
)

func PrintHelloWorld() {
	fmt.Println("Hello World")
}

func GreetUser() {
	var w string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&w)

	fmt.Printf("Hello %s\n", w)
}

func GreetBobOrAlice() {
	var w string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&w)

	if w == "Alice" || w == "Bob" {
		fmt.Printf("Hello %s\n", w)
	} else {
		fmt.Printf("Hello")
	}

}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func AddIntegers(a int, b int) int {
	return a + b
}

func SumDigits(number int) int {
	tmp := number
	s := 0

	for tmp != 0 {
		s += tmp % 10
		tmp /= 10
	}

	return s
}
