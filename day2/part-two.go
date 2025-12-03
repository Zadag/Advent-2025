package main

import (
	"fmt"
	"strconv"
	"strings"
)

func partTwo(entries []string) int64 {
	var solution int64 = 0

	for _, entry := range entries {
		parts := strings.Split(entry, "-")

		lower64, _ := strconv.ParseInt(parts[0], 0, 0)
		upper64, _ := strconv.ParseInt(parts[1], 0, 0)

		upper := int(upper64)
		lower := int(lower64)

		for i := lower; i <= upper; i++ {

			num := int64(i)

			str := strconv.FormatInt(num, 10)
			fmt.Println(str)

			strLen := len(str)

			if strLen <= 1 {
				continue
			}

			foundRepeating := false

			for subLen := 1; subLen < strLen; subLen++ {
				if strLen%subLen == 0 {

					repeated := str[:subLen]

					timesToRepeat := strLen / subLen

					testStr := strings.Repeat(repeated, timesToRepeat)

					if testStr == str {
						solution += num
						foundRepeating = true
						break
					}
				}
			}

			if foundRepeating {
				continue
			}
		}
	}
	return solution
}
