package day11

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	numbers := make([]int, 0)
	nums := strings.Split(input, " ")
	for _, num := range nums {
		if num == "" {
			continue
		}
		number, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	for i := 0; i < 25; i++ {
		newNumbers := make([]int, 0)
		for _, num := range numbers {
			if num == 0 {
				num = 1
				newNumbers = append(newNumbers, num)
				continue
			}

			if len(strconv.Itoa(num))%2 == 0 {
				numString := strconv.Itoa(num)
				firstHalf := numString[:len(numString)/2]
				secondHalf := numString[len(numString)/2:]

				firstHalfInt, err := strconv.Atoi(firstHalf)
				if err != nil {
					panic(err)
				}
				secondHalfInt, err := strconv.Atoi(secondHalf)
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
		numbers = newNumbers
	}

	return strconv.Itoa(len(numbers))
}
