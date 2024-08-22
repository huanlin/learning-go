package main

import "fmt"

func main() {
	// Use hard-coded text to start.
	text := "let's count some words"

	var numSpaces int

	// Look at the text one byte at a time.
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			numSpaces++
		}
	}

	// The number of words is one more than the number of spaces.
	numWords := numSpaces + 1
	fmt.Printf("There are %d words in the text.\n", numWords)
}
