package day3

import (
	"regexp"
	"strconv"
)

func Part1(input string) string {
	match, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := match.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, m := range matches {
		l, _ := strconv.Atoi(m[1])
		r, _ := strconv.Atoi(m[2])

		sum += l * r

	}

	return strconv.Itoa(sum)
}
