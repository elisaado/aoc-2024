package day2

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input string) string {
	splat := strings.Split(input, "\n")
	levels := make([][]int, 0)
	for _, s := range splat {
		q := strings.Split(s, " ")
		level := make([]int, 0)
		for _, r := range q {
			if r == "" {
				continue
			}
			elem, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}
			level = append(level, elem)
		}
		levels = append(levels, level)
	}

	safeLevels := 0
	for _, level := range levels {
		shouldBeIncreasing := true
		for i, elem := range level {
			if i == 0 {
				if elem >= level[i+1] {
					shouldBeIncreasing = false
				}
				continue
			}

			diff := elem - level[i-1]
			if diff < 0 {
				if shouldBeIncreasing {
					break
				}
			}
			if diff > 0 {
				if !shouldBeIncreasing {
					break
				}
			}

			diff = int(math.Abs(float64(diff)))
			if diff == 0 {
				break
			}
			if diff > 3 {
				break
			}
			if i == len(level)-1 {
				safeLevels += 1
			}
		}
	}
	return strconv.Itoa(safeLevels)
}
