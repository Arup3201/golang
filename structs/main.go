package main

import (
	"fmt"
	"golang/structs/person"
)

func main() {
	p := person.NewPerson("Arup", 23, "Kolkata")

	fmt.Printf("Name: %s, Age: %d, Location: %s\n", p.Name, p.Age, p.Location)

	fmt.Println(p.SayHello("Subha"))
	fmt.Println(p.CelebrateBirthday())

	fmt.Printf("Name: %s, Age: %d, Location: %s\n", p.Name, p.Age, p.Location)
}
