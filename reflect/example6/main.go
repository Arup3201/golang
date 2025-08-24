package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name     string
	Location string
}

func main() {
	p := Person{
		Name:     "John Dao",
		Location: "India",
	}

	s := reflect.ValueOf(&p).Elem()
	fmt.Printf("s = {Name: %s, Location: %s}", s.Field(0).Interface(), s.Field(1).Interface())

	s.Field(0).SetString("Arup")
	s.Field(1).SetString("Tamluk")
	fmt.Printf("s = {Name: %s, Location: %s}", s.Field(0).Interface(), s.Field(1).Interface())
}
