package day3

import (
	"regexp"
	"strconv"
)

func Part2(input string) string {
	match, _ := regexp.Compile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)
	matches := match.FindAllStringSubmatch(input, -1)

	sum := 0
	doing := true
	for _, m := range matches {
		for _, n := range m {
			if n == "do()" {
				doing = true
			}

			if n == "don't()" {
				doing = false
			}
		}

		l, err := strconv.Atoi(m[1])
		if err != nil {
			continue
		}
		r, err := strconv.Atoi(m[2])
		if err != nil {
			continue
		}

		if doing {
			sum += l * r
		}

	}

	return strconv.Itoa(sum)
}
