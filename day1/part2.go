package day1

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
	left := make([]int, 0)
	right := make([]int, 0)
	splat := strings.Split(input, "\n")
	for _, s := range splat {
		if s == "" {
			continue
		}
		leftAndRight := strings.Split(s, " ")

		leftElem, err := strconv.Atoi(strings.TrimSpace(leftAndRight[0]))
		if err != nil {
			panic(err)
		}

		rightElem, err := strconv.Atoi(strings.TrimSpace(leftAndRight[len(leftAndRight)-1]))
		if err != nil {
			panic(err)
		}

		left = append(left, leftElem)
		right = append(right, rightElem)
	}

	sum := 0
	for _, l := range left {
		seen := 0
		for _, r := range right {
			if l == r {
				seen += 1
			}
		}

		sum += l * seen
	}

	return strconv.Itoa(sum)
}
