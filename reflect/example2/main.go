package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type()) // same as reflect.TypeOf(x)
	fmt.Println("kind:", v.Kind())
	fmt.Println("kind is unint8:", v.Kind() == reflect.Uint8)
	fmt.Println("getter operates on type:", reflect.TypeOf(v.Uint())) // getter and setter operates on the largest type that can hold the value
}
