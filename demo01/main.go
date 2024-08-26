/*
Demo01 is a simple program that prints "Hello, World!" to the console.
It also demonstrates how to write tests for Go programs.

(1) To build the program, navigate to the directory containing the main.go file and run:

go build

This will compile the program and create an executable file in the same directory.

(2) To run the program, navigate to the directory containing the executable file and run:

./demo01

This will run the program and print "Hello, World!" to the console.

(3) To run the tests, navigate to the directory containing the hello_test.go file and run:

go test

This will run the tests and print the results to the console.
*/
package main

import "fmt"

func main() {
	fmt.Println(hello())
}
