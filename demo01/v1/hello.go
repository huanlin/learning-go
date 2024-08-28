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

	a := Animal{
		name: "cat",
	}

	fmt.Println(a.speak())
	fmt.Println(a.name)
}
