package main

import (
	"fmt"
)

var case1 bool = true
var sum int = 100

func main() {
	if case1 {
		sum := add(5, 5) // 區域變數
		fmt.Println(sum)
	} else {
		sum := add(10, 10) // 區域變數
		fmt.Println(sum)
	}

	fmt.Println(sum) // 使用全域變數
}

func add(x, y int) int {
	return x + y
}
