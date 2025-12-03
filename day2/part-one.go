package main

import (
	"fmt"
	"strconv"
	"strings"
)

func partOne(entries []string) {
	solution := 0
	for _, entry := range entries {
		lower, _ := strconv.ParseInt(strings.Split(entry, "-")[0], 0, 0)
		upper, _ := strconv.ParseInt(strings.Split(entry, "-")[1], 0, 0)

		fmt.Println(lower, upper)

		for i := lower; i <= upper; i++ {
			if isTwoOfTheSame(strconv.FormatInt(i, 10)) {
				solution += int(i)
			}
		}
	}
	fmt.Println(solution)
}

func isTwoOfTheSame(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(s)/2:]

	if firstHalf == secondHalf {
		return true
	}

	fmt.Println("first and second", firstHalf, secondHalf)
	return false
}
