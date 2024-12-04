package day4

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
	var grid [][]string
	splat := strings.Split(input, "\n")
	for _, row := range splat {
		if row == "" {
			continue
		}
		grid = append(grid, strings.Split(row, ""))
	}

	stringsInGrid := make([]string, 0)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if y < len(grid)-2 && x < len(grid[y])-2 {
				stringInGrid := grid[y][x] + grid[y+1][x+1] + grid[y+2][x+2] + grid[y][x+2] + grid[y+1][x+1] + grid[y+2][x]
				stringsInGrid = append(stringsInGrid, stringInGrid)
			}
		}
	}

	sum := 0
	for _, str := range stringsInGrid {
		if str == "MASSAM" {
			sum += 1
		}

		if str == "MASMAS" {
			sum += 1
		}

		if str == "SAMMAS" {
			sum += 1
		}

		if str == "SAMSAM" {
			sum += 1
		}
	}

	return strconv.Itoa(sum)
}
