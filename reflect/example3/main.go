package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func main() {
	var i MyInt
	var j int

	fmt.Println("Type of i:", reflect.TypeOf(i)) // MyInt
	fmt.Println("Type of j:", reflect.TypeOf(j))

	// kind holds the underlying type, not the static type
	// kind can't discrimate between MyInt and int even though Type can
	fmt.Println("kind of i:", reflect.ValueOf(i).Kind()) // int
	fmt.Println("kind of j:", reflect.ValueOf(j).Kind())
}
