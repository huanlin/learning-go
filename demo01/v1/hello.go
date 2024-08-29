package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a *Animal) speak() string {
	a.name = "aa"
	switch a.name {
	case "cat":
		return "meow"
	case "dog":
		return "woof"
	default:
		return "nondescript animal noise?"
	}
}

func main() {
	fmt.Println("Hello, World!")

	cat1 := new(Animal)
	cat2 := &Animal{}
	var cat3 *Animal = new(Animal)

	cat1.name = "cat"
	cat2.name = "cat"

	fmt.Printf("%T\n", cat1)
	fmt.Printf("%T\n", cat2)
	fmt.Printf("%T\n", cat3)

	var x interface{} = 7

	switch x.(type) {
	case int:
		fmt.Println("int")
	}
}
