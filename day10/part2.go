package day10

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
	grid := make([][]Point, 0)
	lines := strings.Split(input, "\n")
	trailheads := make([]Point, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]Point, 0)
		for _, char := range line {
			if char != '.' {
				parsed, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}
				point := Point{
					X:         len(row),
					Y:         len(grid),
					Height:    parsed,
					Trailhead: parsed == 0,
				}
				row = append(row, point)
				if point.Trailhead {
					trailheads = append(trailheads, point)
				}
			}
		}
		grid = append(grid, row)
	}

	sum := 0
	for _, trailhead := range trailheads {
		sum += ReachableTrailPaths(grid, trailhead, make([]Point, 0))
	}

	return strconv.Itoa(sum)
}

func ReachableTrailPaths(grid [][]Point, start Point, visited []Point) int {
	for _, visit := range visited {
		if visit.X == start.X && visit.Y == start.Y {
			return 0
		}
	}

	visited = append(visited, start)
	if start.Height == 9 {
		return 1
	}

	neighbours := make([]Point, 0)
	if start.X > 0 {
		neighbours = append(neighbours, grid[start.Y][start.X-1])
	}
	if start.X < len(grid[0])-1 {
		neighbours = append(neighbours, grid[start.Y][start.X+1])
	}
	if start.Y > 0 {
		neighbours = append(neighbours, grid[start.Y-1][start.X])
	}
	if start.Y < len(grid)-1 {
		neighbours = append(neighbours, grid[start.Y+1][start.X])
	}

	reachable := 0
	for _, neighbour := range neighbours {
		if neighbour.Height-start.Height == 1 {
			reachable += ReachableTrailPaths(grid, neighbour, visited)
		}
	}

	return reachable
}
