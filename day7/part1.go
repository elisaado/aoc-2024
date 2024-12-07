package day7

import (
	"math"
	"strconv"
	"strings"
)

type Operator int

const (
	None              = -1
	Add      Operator = 0
	Multiply Operator = 1

	// part 2
	Concat Operator = 2
)

type Eq struct {
	Left     int
	Right    *Eq
	Operator Operator
}

func (e *Eq) EvaluateIteratively() int {
	if e.Right == nil {
		return e.Left
	}

	result := e.Left
	for e.Right != nil {
		if e.Operator == Add {
			result += e.Right.Left
		}

		if e.Operator == Multiply {
			result *= e.Right.Left
		}

		// part 2
		if e.Operator == Concat {
			// result = result*10 + e.Right.Left <-- DOES NOT WORK (test using 12 || 345)
			result = result*int(math.Pow(10, float64(len(strconv.Itoa(e.Right.Left))))) + e.Right.Left // <-- does work
		}

		e = e.Right
	}

	return result
}

func (e *Eq) String() string {
	if e.Right == nil {
		return strconv.Itoa(e.Left)
	}

	var opString string
	if e.Operator == Add {
		opString = "+"
	}
	if e.Operator == Multiply {
		opString = "*"
	}

	// part 2
	if e.Operator == Concat {
		opString = "||"
	}

	return strconv.Itoa(e.Left) + " " + opString + " " + e.Right.String()
}

func (e *Eq) Length() int {
	len := 0
	for e != nil {
		len++
		e = e.Right
	}
	return len
}

func Part1(input string) string {
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
		for i := 0; i < int(math.Pow((2), length)); i++ {
			changeAbleEq := equations[key]
			bits := i
			for changeAbleEq.Right != nil {
				bit := bits & 1
				changeAbleEq.Operator = Operator(bit)
				bits >>= 1
				changeAbleEq = changeAbleEq.Right
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
