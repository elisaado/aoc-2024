package day13

import (
	"regexp"
	"strconv"
	"strings"
)

func Part2(input string) string {
	// Parse input
	prizeLines := strings.Split(input, "\n\n")

	prizes := make([]Prize, 0)
	for _, prizeLine := range prizeLines {
		if prizeLine == "" {
			continue
		}
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
		prize.X = ToInt(matches[2]) + 10000000000000
		prize.Y = ToInt(matches[4]) + 10000000000000

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

		// we use cramers rule to solve the system of equations
		// https://en.wikipedia.org/wiki/Cramer%27s_rule
		m := prize.X*prize.ButtonBy - prize.ButtonBx*prize.Y
		n := prize.ButtonAx*prize.ButtonBy - prize.ButtonBx*prize.ButtonAy
		if m%n != 0 {
			println("cant win prize at ", prize.X, prize.Y)
			continue
		}
		p := m / n

		m = prize.ButtonAx*prize.Y - prize.X*prize.ButtonAy
		n = prize.ButtonAx*prize.ButtonBy - prize.ButtonBx*prize.ButtonAy
		if m%n != 0 {
			println("cant win prize at ", prize.X, prize.Y)
			continue
		}
		q := m / n

		println(prize.X)
		println(p, q)

		if p < 0 || q < 0 {
			println("cant win prize at ", prize.X, prize.Y)
			continue
		}

		// y part
		sum += p*3 + q
	}

	return strconv.Itoa(sum)
}
