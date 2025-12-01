package main

import (
	"fmt"
	"os"
)

func main() {
	// Get a a string
	inputStr, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("%s: Error reading file", err)
		return
	}

	//partOne(inputStr)
	partTwo(inputStr)
}
