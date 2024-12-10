package day10

import (
	"strconv"
	"strings"
)

type Point struct {
	X         int
	Y         int
	Height    int
	Trailhead bool
}

func Part1(input string) string {
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
		rechable := ReachableTrailends(grid, trailhead, make([]Point, 0))
		sum += len(rechable)
	}

	return strconv.Itoa(sum)
}

func ReachableTrailends(grid [][]Point, start Point, visited []Point) []Point {
	for _, visit := range visited {
		if visit.X == start.X && visit.Y == start.Y {
			return make([]Point, 0)
		}
	}

	visited = append(visited, start)
	if start.Height == 9 {
		return []Point{start}
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

	reachable := make([]Point, 0)
	for _, neighbour := range neighbours {
		if neighbour.Height-start.Height == 1 {
			reachable = append(reachable, ReachableTrailends(grid, neighbour, visited)...)
		}
	}

	// remove duplicates
	reachableMap := make(map[Point]bool)
	for _, point := range reachable {
		reachableMap[point] = true
	}

	reachable = make([]Point, 0)
	for point := range reachableMap {
		reachable = append(reachable, point)
	}

	return reachable
}
