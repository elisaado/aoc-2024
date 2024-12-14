package day13

import (
	"regexp"
	"strconv"
	"strings"
)

type Prize struct {
	X        int
	Y        int
	ButtonAx int
	ButtonAy int
	ButtonBx int
	ButtonBy int
}

func Part1(input string) string {
	// Parse input
	prizeLines := strings.Split(input, "\n\n")

	prizes := make([]Prize, 0)
	for _, prizeLine := range prizeLines {
		line := strings.Split(prizeLine, "\n")
		var prize Prize
		re := regexp.MustCompile(`X(\+|=)(\d+),\sY(\+|=)(\d+)`)
		matches := re.FindStringSubmatch(line[0])
		prize.ButtonAx = ToInt(matches[2])
		prize.ButtonAy = ToInt(matches[4])

		matches = re.FindStringSubmatch(line[1])
		prize.ButtonBx = ToInt(matches[2])
		prize.ButtonBy = ToInt(matches[4])

		matches = re.FindStringSubmatch(line[2])
		prize.X = ToInt(matches[2])
		prize.Y = ToInt(matches[4])

		prizes = append(prizes, prize)
	}

	sum := 0
	for _, prize := range prizes {
		// x part
		if GCD(prize.ButtonAx, prize.ButtonBx) > 1 {
			if prize.X%GCD(prize.ButtonAx, prize.ButtonBx) != 0 {
				println("cant win prize at ", prize.X, prize.Y)
				continue
			}
		}

		if GCD(prize.ButtonAy, prize.ButtonBy) > 1 {
			if prize.Y%GCD(prize.ButtonAy, prize.ButtonBy) != 0 {
				println("(y case) cant win prize at ", prize.X, prize.Y)
				continue
			}
		}

		p := 0
		q := 0
		cost := 9999999999
		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				xVal := prize.ButtonAx*i + prize.ButtonBx*j
				yVal := prize.ButtonAy*i + prize.ButtonBy*j
				if xVal > prize.X || yVal > prize.Y {
					break
				}

				if xVal == prize.X && yVal == prize.Y && cost > i*3+j {
					p = i
					q = j
					cost = p*3 + q
				}
			}
		}

		if p == 0 && q == 0 {
			println("cant win prize at ", prize.X, prize.Y)
			continue
		}

		println(p, q)

		// y part
		sum += p*3 + q
	}

	return strconv.Itoa(sum)
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	if a == 0 {
		println(a, b)
		panic("GCD is 0")
	}

	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
