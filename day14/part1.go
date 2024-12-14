package day14

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func Part1(input string) string {
	test := false
	if len(os.Args) > 4 {
		test = true
	}
	gridSizeX := 101
	gridSizeY := 103
	if test {
		gridSizeX = 11
		gridSizeY = 7
	}

	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	lines := strings.Split((input), "\n")
	robots := make([]Robot, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := pattern.FindStringSubmatch(line)
		robot := Robot{
			x:  ToInt(matches[1]),
			y:  ToInt(matches[2]),
			vx: ToInt(matches[3]),
			vy: ToInt(matches[4]),
		}
		robots = append(robots, robot)
	}

	seconds := 100
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for _, robot := range robots {
		robot.x = (robot.x + robot.vx*seconds) % gridSizeX
		if robot.x < 0 {
			robot.x += gridSizeX
		}

		robot.y = (robot.y + robot.vy*seconds) % gridSizeY
		if robot.y < 0 {
			robot.y += gridSizeY
		}

		if robot.x < gridSizeX/2 && robot.y < gridSizeY/2 {
			q1++
		}

		if robot.x > gridSizeX/2 && robot.y < gridSizeY/2 {
			q2++
		}

		if robot.x < gridSizeX/2 && robot.y > gridSizeY/2 {
			q3++
		}

		if robot.x > gridSizeX/2 && robot.y > gridSizeY/2 {
			q4++
		}
	}

	return strconv.Itoa(q1 * q2 * q3 * q4)
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
