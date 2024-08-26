package main

import (
	"flag"
	"fmt"
)

func hello(name string) {
	fmt.Printf("Hello, %s!", name)
}

var name string

func init() {
	name = *flag.String("name", "World", "a name to say hello to")
}

func main() {
	hello(name)
}
