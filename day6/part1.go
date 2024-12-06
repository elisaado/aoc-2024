package day6

import (
	"fmt"
	"strconv"
	"strings"
)

// might implement this later instead of using stringd
// const (
// 	Start int = iota
// 	Visited int

// )

func Part1(input string) string {
	grid := make([][]string, 0)

	splat := strings.Split(input, "\n")
	startCoords := make([]int, 2)
	direction := ""

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
				direction = char
			}

			row = append(row, char)
		}

		grid = append(grid, row)
	}

	currentCoords := startCoords
	encounteredObstacles := 0

	for {
		oldCoords := currentCoords
		bounced := false
		if direction == "^" {
			for i := currentCoords[1]; i >= 0; i-- {
				if grid[i][currentCoords[0]] == "#" {
					encounteredObstacles++
					currentCoords = []int{currentCoords[0], i + 1}
					direction = ">"
					bounced = true
					break
				}
			}

			if !bounced {
				currentCoords = []int{currentCoords[0], 0}
			}

			for i := oldCoords[1]; i >= currentCoords[1]; i-- {
				if grid[i][currentCoords[0]] != "#" {
					grid[i][currentCoords[0]] = "X"
				}
			}
		} else if direction == "v" {
			for i := currentCoords[1]; i < len(grid); i++ {
				if grid[i][currentCoords[0]] == "#" {
					encounteredObstacles++
					currentCoords = []int{currentCoords[0], i - 1}
					direction = "<"
					bounced = true
					break
				}
			}

			if !bounced {
				currentCoords = []int{currentCoords[0], len(grid) - 1}
			}

			for i := oldCoords[1]; i <= currentCoords[1]; i++ {
				if grid[i][currentCoords[0]] != "#" {
					grid[i][currentCoords[0]] = "X"
				}
			}
		} else if direction == "<" {
			for i := currentCoords[0]; i >= 0; i-- {
				if grid[currentCoords[1]][i] == "#" {
					encounteredObstacles++
					currentCoords = []int{i + 1, currentCoords[1]}
					direction = "^"
					bounced = true
					break
				}
			}

			if !bounced {
				currentCoords = []int{0, currentCoords[1]}
			}

			for i := oldCoords[0]; i >= currentCoords[0]; i-- {
				if grid[currentCoords[1]][i] != "#" {
					grid[currentCoords[1]][i] = "X"
				}
			}
		} else if direction == ">" {
			for i := currentCoords[0]; i < len(grid[currentCoords[1]]); i++ {
				if grid[currentCoords[1]][i] == "#" {
					encounteredObstacles++
					currentCoords = []int{i - 1, currentCoords[1]}
					direction = "v"
					bounced = true
					break
				}
			}

			if !bounced {
				currentCoords = []int{len(grid[currentCoords[1]]) - 1, currentCoords[1]}
			}

			for i := oldCoords[0]; i <= currentCoords[0]; i++ {
				if grid[currentCoords[1]][i] != "#" {
					grid[currentCoords[1]][i] = "X"
				}
			}
		}

		fmt.Println("after iteration")
		fmt.Println(currentCoords[0], currentCoords[1])

		if !bounced {
			println("done!")
			println(direction)
			break
		}
	}

	// pretty print grid

	for _, row := range grid {
		for _, char := range row {
			print(char)
		}
		println()
	}

	count := 0
	for _, row := range grid {
		for _, char := range row {
			if char == "X" {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}
