package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	v := reflect.ValueOf(x)

	fmt.Println("settability of v:", v.CanSet())

	// will panic as v is not settable
	// v.SetFloat(5.22)

	p := reflect.ValueOf(&v)                     // pass address instead of value
	fmt.Println("settability of p:", p.CanSet()) // still false, but that is for p, we want to modify *p

	// get the Elem() which indirects through the pointer
	s := p.Elem()
	fmt.Println("settability of s:", s.CanSet()) // true
}
