package day4

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	var grid [][]string
	splat := strings.Split(input, "\n")
	for _, row := range splat {
		if row == "" {
			continue
		}
		grid = append(grid, strings.Split(row, ""))
	}

	sum := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "X" {
				// Check left
				if y > 2 && grid[y-1][x] == "M" {
					if grid[y-2][x] == "A" {
						if grid[y-3][x] == "S" {
							sum += 1
						}
					}
				}

				// Check right
				if y < len(grid)-3 && grid[y+1][x] == "M" {
					if grid[y+2][x] == "A" {
						if grid[y+3][x] == "S" {
							sum += 1
						}
					}
				}

				// Check up
				if x > 2 && grid[y][x-1] == "M" {
					if grid[y][x-2] == "A" {
						if grid[y][x-3] == "S" {
							sum += 1
						}
					}
				}

				// Check down
				if x < len(grid[y])-3 && grid[y][x+1] == "M" {
					if grid[y][x+2] == "A" {
						if grid[y][x+3] == "S" {
							sum += 1
						}
					}
				}

				// Check diagonal up left
				if y > 2 && x > 2 && grid[y-1][x-1] == "M" {
					if grid[y-2][x-2] == "A" {
						if grid[y-3][x-3] == "S" {
							sum += 1
						}
					}
				}

				// Check diagonal up right
				if y > 2 && x < len(grid[y])-3 && grid[y-1][x+1] == "M" {
					if grid[y-2][x+2] == "A" {
						if grid[y-3][x+3] == "S" {
							sum += 1
						}
					}
				}

				// Check diagonal down left
				if y < len(grid)-3 && x > 2 && grid[y+1][x-1] == "M" {
					if grid[y+2][x-2] == "A" {
						if grid[y+3][x-3] == "S" {
							sum += 1
						}
					}
				}

				// Check diagonal down right
				if y < len(grid)-3 && x < len(grid[y])-3 && grid[y+1][x+1] == "M" {
					if grid[y+2][x+2] == "A" {
						if grid[y+3][x+3] == "S" {
							sum += 1
						}
					}
				}

			}
		}
	}

	return strconv.Itoa(sum)
}
