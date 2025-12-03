package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get a a string
	inputStr, err := os.ReadFile("part1-input.txt")
	if err != nil {
		fmt.Printf("%s: Error reading file", err)
		return
	}

	entries := strings.Split(string(inputStr), ",")

	//partOne(entries)
	partTwo(entries)

}
