package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part2(input string) string {
	splat := strings.Split(input, "\n")
	equations := make(map[int]*Eq, 0)
	for _, line := range splat {
		if len(line) == 0 {
			continue
		}

		equation := strings.Split(line, ":")
		result, err := strconv.Atoi(equation[0])
		if err != nil {
			panic(err)
		}
		equation = strings.Split(strings.TrimSpace(equation[1]), " ")

		var eqs []*Eq

		for _, char := range equation {
			var eq Eq = Eq{}
			if char == "" {
				continue
			}
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}

			eq.Left = num
			eq.Operator = Multiply
			eq.Right = nil

			eqs = append(eqs, &eq)
		}

		// doing it like this because Go likes to pass by reference(?)
		for i, eq := range eqs {
			if i+1 < len(eqs) {
				eq.Right = eqs[i+1]
			} else {
				eq.Right = nil
			}
		}

		equations[result] = eqs[0]
	}

	canBeSolved := make([]int, 0)
	for key := range equations {
		eq := equations[key]
		length := float64(eq.Length() - 1)
		for i := 0; i < int(math.Pow((3), length)); i++ {
			changeAbleEq := equations[key]
			ding := fmt.Sprintf("%0*s", int(length), strconv.FormatInt(int64(i), 3))

			for ding != "" {
				lastNum, err := strconv.Atoi(string(ding[len(ding)-1]))
				if err != nil {
					panic(err)
				}
				changeAbleEq.Operator = Operator(lastNum)
				changeAbleEq = changeAbleEq.Right

				ding = ding[:len(ding)-1]
			}
			if eq.EvaluateIteratively() == key {
				canBeSolved = append(canBeSolved, key)
				break
			}
		}
	}

	sum := 0
	for _, m := range canBeSolved {
		sum += m
	}

	return strconv.Itoa(sum)
}
