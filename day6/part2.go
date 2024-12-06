package day6

import (
	"strconv"
	"strings"
)

// brute force solution, but it works reasonably fast!
// lots of code duplication because of early returns so it runs faster
// thanks for listening to my TED talk

type Location struct {
	X         int
	Y         int
	direction string
}

func Part2(input string) string {
	grid := make([][]string, 0)

	splat := strings.Split(input, "\n")
	startCoords := make([]int, 2)

	for _, line := range splat {
		if len(line) == 0 {
			continue
		}
		chars := strings.Split(line, "")
		row := make([]string, 0)

		for _, char := range chars {
			if char == "" {
				continue
			}

			if char == "^" || char == "v" || char == "<" || char == ">" {
				startCoords[0] = len(row)
				startCoords[1] = len(grid)
			}

			row = append(row, char)
		}

		grid = append(grid, row)
	}

	numOfLoops := 0

	for y, row := range grid {
		for x, char := range row {
			// deep copy the grid
			gridCopy := make([][]string, len(grid))
			for i := range grid {
				gridCopy[i] = make([]string, len(grid[i]))
				copy(gridCopy[i], grid[i])
			}
			if char == "." {
				gridCopy[y][x] = "#"
				if checkLoop(gridCopy) {
					numOfLoops++
				}
			}
		}
	}

	return strconv.Itoa(numOfLoops)
}

func checkLoop(grid [][]string) bool {
	startCoords := make([]int, 2)
	direction := ""

	for y, row := range grid {
		for x, char := range row {
			if char == "^" || char == "v" || char == "<" || char == ">" {
				startCoords[0] = x
				startCoords[1] = y
				direction = char
			}
		}
	}

	currentCoords := startCoords

	visitedLocations := make([]Location, 0)
	loop := false
	for {
		bounced := false
		if direction == "^" {
			for i := currentCoords[1]; i >= 0; i-- {
				if grid[i][currentCoords[0]] == "#" {
					currentCoords = []int{currentCoords[0], i + 1}
					direction = ">"
					bounced = true
					// if we've been here before, looking in this direction, we're in a loop
					loop = checkAndAppendLocation(&visitedLocations, currentCoords, direction)
					if loop {
						return true
					}
					break
				}
			}

			if !bounced {
				currentCoords = []int{currentCoords[0], 0}
			}
		} else if direction == "v" {
			for i := currentCoords[1]; i < len(grid); i++ {
				if grid[i][currentCoords[0]] == "#" {
					currentCoords = []int{currentCoords[0], i - 1}
					direction = "<"
					bounced = true
					loop = checkAndAppendLocation(&visitedLocations, currentCoords, direction)
					if loop {
						return true
					}
					break
				}
			}

			if !bounced {
				currentCoords = []int{currentCoords[0], len(grid) - 1}
			}
		} else if direction == "<" {
			for i := currentCoords[0]; i >= 0; i-- {
				if grid[currentCoords[1]][i] == "#" {
					currentCoords = []int{i + 1, currentCoords[1]}
					direction = "^"
					bounced = true
					loop = checkAndAppendLocation(&visitedLocations, currentCoords, direction)
					if loop {
						return true
					}
					break
				}
			}

			if !bounced {
				currentCoords = []int{0, currentCoords[1]}
			}
		} else if direction == ">" {
			for i := currentCoords[0]; i < len(grid[currentCoords[1]]); i++ {
				if grid[currentCoords[1]][i] == "#" {
					currentCoords = []int{i - 1, currentCoords[1]}
					direction = "v"
					bounced = true
					loop = checkAndAppendLocation(&visitedLocations, currentCoords, direction)
					if loop {
						return true
					}
					break
				}
			}

			if !bounced {
				currentCoords = []int{len(grid[currentCoords[1]]) - 1, currentCoords[1]}
			}
		}

		if !bounced {
			return false
		}

		if loop {
			panic("we shouldn't be here")
		}
	}
}

func checkAndAppendLocation(locations *[]Location, coords []int, direction string) bool {
	for _, location := range *locations {
		if location.X == coords[0] && location.Y == coords[1] && location.direction == direction {
			return true
		}
	}

	*locations = append(*locations, Location{X: coords[0], Y: coords[1], direction: direction})
	return false
}
