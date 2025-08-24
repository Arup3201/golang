package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := reflect.ValueOf(3.14)
	fmt.Println("Interface value:", x.Interface().(float64)) // going back from a reflection object of interface, and type asserting it using float64 (optional)
}
