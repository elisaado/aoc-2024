package day14

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2(input string) string {
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

	// brute force :DD
	seconds := 10000000

	neededSeconds := 0
	for i := 0; i < seconds; i++ {
		grid := make([][]int, gridSizeY)
		for i := range grid {
			grid[i] = make([]int, gridSizeX)
		}

		for _, robot := range robots {
			x := (robot.x + robot.vx*i) % gridSizeX
			if x < 0 {
				x += gridSizeX
			}

			y := (robot.y + robot.vy*i) % gridSizeY
			if y < 0 {
				y += gridSizeY
			}

			grid[y][x] = 1
		}

		robotsNextToEachOther := 0
		for y := 0; y < gridSizeY; y++ {
			for x := 0; x < gridSizeX-1; x++ {
				if grid[y][x] == 1 {
					if grid[y][x+1] == 1 {
						robotsNextToEachOther++
					}
				}
			}
		}

		// just a guess of how many robots
		// need to be next to each other
		// to print a christmas tree
		if robotsNextToEachOther > 100 {
			// pretty print the grid
			neededSeconds = i
			break
		}
	}

	return strconv.Itoa(neededSeconds)
}
