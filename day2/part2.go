package day2

import (
	"math"
	"strconv"
	"strings"
)

func Part2(input string) string {
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
	// brute force yippie
	for _, level := range levels {
		if checkSafe(level) {
			safeLevels += 1
			continue
		}
		for i := 0; i < len(level); i++ {
			if checkSafe(excludeElement(level, i)) {
				safeLevels += 1
				break
			}
		}
	}

	return strconv.Itoa(safeLevels)
}

func excludeElement(level []int, index int) []int {
	newLevel := make([]int, 0)
	for i, elem := range level {
		if i == index {
			continue
		}
		newLevel = append(newLevel, elem)
	}
	return newLevel
}

func checkSafe(level []int) bool {
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
				return false
			}
		}
		if diff > 0 {
			if !shouldBeIncreasing {
				return false
			}
		}

		diff = int(math.Abs(float64(diff)))
		if diff == 0 {
			return false
		}
		if diff > 3 {
			return false
		}
		if i == len(level)-1 {
			return true
		}
	}
	return false
}
