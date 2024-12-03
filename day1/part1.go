package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) string {
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

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sum := 0

	for i, l := range left {

		r := right[i]
		sum += int(math.Abs(float64(r - l)))

	}

	return strconv.Itoa(sum)
}
