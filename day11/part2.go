package day11

// dit is niet hoe het moet
// https://gitlab.com/14mRh4X0r/aoc-2024/-/blob/master/11b.rb?ref_type=heads <- is wel hoe het moet

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
	numbers := make([]int64, 0)
	nums := strings.Split(input, " ")
	for _, num := range nums {
		if num == "" {
			continue
		}
		number, err := strconv.ParseInt(strings.TrimSpace(num), 10, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	counts := make(map[int64]int64, 0)
	for _, num := range numbers {
		counts[num]++
	}

	for i := 0; i < 75; i++ {
		newCounts := make(map[int64]int64, 0)

		for num, count := range counts {
			ints := ExpandNumber(num)

			countsForInts := make(map[int64]int64, 0)
			for _, newNum := range ints {
				countsForInts[newNum]++
			}

			for newNum, newCount := range countsForInts {
				newCounts[newNum] += newCount * count
			}
		}

		counts = newCounts
	}

	var totalNums int64 = 0
	for _, count := range counts {
		totalNums += count
	}

	return strconv.FormatInt(totalNums, 10)
}

func ExpandNumber(num int64) []int64 {
	numbers := make([]int64, 0)
	numbers = append(numbers, num)

	newNumbers := make([]int64, 0)
	for _, num := range numbers {
		if num == 0 {
			num = 1
			newNumbers = append(newNumbers, num)
			continue
		}

		if len(strconv.FormatInt(num, 10))%2 == 0 {
			numString := strconv.FormatInt(num, 10)
			firstHalf := numString[:len(numString)/2]
			secondHalf := numString[len(numString)/2:]

			firstHalfInt, err := strconv.ParseInt(firstHalf, 10, 64)
			if err != nil {
				panic(err)
			}
			secondHalfInt, err := strconv.ParseInt(secondHalf, 10, 64)
			if err != nil {
				panic(err)
			}
			newNumbers = append(newNumbers, firstHalfInt)
			newNumbers = append(newNumbers, secondHalfInt)
			continue
		}

		num = num * 2024
		newNumbers = append(newNumbers, num)
	}

	return newNumbers
}
