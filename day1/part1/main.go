package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputStr, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("%s: Error reading file", err)
		return
	}

	ops := strings.Split(string(inputStr), "\n")

	pos := 50
	solution := 0
	for _, val := range ops {
		direction := string(val[0])
		distance, _ := strconv.ParseInt(val[1:], 0, 64)

		if direction == "L" {
			pos -= int(distance)
		}
		if direction == "R" {
			pos += int(distance)
		}
		fmt.Println(val, distance, direction, pos%99)
		if pos%100 == 0 {
			solution++
		}
	}
	fmt.Println(solution)
}
