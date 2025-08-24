package main

import (
	"fmt"
	"reflect"
)

/*
	{}interface -> (value, type)
	reflect.TypeOf() -> gives the type description of the interface value
	reflect.ValueOf() -> gives the concrete value
*/

func main() {
	var x float64 = 3.14
	fmt.Println("type of (x):", reflect.TypeOf(x))
	fmt.Println("value of (x):", reflect.ValueOf(x).String())
}
