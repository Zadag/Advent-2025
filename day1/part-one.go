package main

import (
	"fmt"
	"strconv"
	"strings"
)

func partOne(str []byte) {
	ops := strings.Split(string(str), "\n")

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
		// fmt.Println(val, distance, direction, pos%100)
		if pos%100 == 0 {
			solution++
		}
	}
	fmt.Println(solution)
}
