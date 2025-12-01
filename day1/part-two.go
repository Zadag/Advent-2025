package main

import (
	"fmt"
	"strconv"
	"strings"
)

func partTwo(str []byte) {
	ops := strings.Split(string(str), "\n")

	pos := 50
	solution := 0

	for _, val := range ops {
		direction := string(val[0])
		distance, _ := strconv.ParseInt(val[1:], 0, 64)

		for mag := range distance {
			fmt.Println(mag) // dumb log, go compiler complains if mag is unused

			if direction == "R" {
				pos++
				if pos%100 == 0 {
					solution++
				}
			} else {
				pos--
				if pos%100 == 0 {
					solution++
				}
			}
		}

	}

	fmt.Println(solution)
}
