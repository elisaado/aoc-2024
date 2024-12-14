package day12

import (
	"slices"
	"strconv"
	"strings"
)

func Part2(input string) string {
	grid := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		grid = append(grid, strings.Split(line, ""))
	}

	totalVisited := make([][]int, 0)
	letters := make(map[string][][]int)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			alreadyVisited := slices.ContainsFunc(totalVisited, func(v []int) bool {
				return v[0] == x && v[1] == y
			})
			for _, v := range totalVisited {
				if v[0] == x && v[1] == y {
					alreadyVisited = true
					break
				}
			}

			if alreadyVisited {
				println("Already visited", grid[y][x])
				continue
			}

			area, visited := GetArea(grid, x, y)
			circum := GetNumEdges(grid, x, y)
			if area == 0 {
				continue
			}

			c := grid[y][x]
			if letters[c] == nil {
				letters[c] = make([][]int, 0)
			}
			letters[c] = append(letters[c], []int{area, circum})

			for _, v := range visited {
				c2 := grid[v[1]][v[0]]
				if c != c2 {
					panic("Different letters in area???")
				}
			}
			totalVisited = append(totalVisited, visited...)
		}
	}

	sum := 0

	for c, v := range letters {
		for _, vv := range v {
			println("Letter: ", c)
			println("Letter area:", vv[0], "Letter edges:", vv[1])
			sum += vv[0] * vv[1]
		}
	}

	return strconv.Itoa(sum)
}

func GetNumEdges(grid [][]string, x, y int) int {
	corners := 0
	c := grid[y][x]
	cellsInGroup, _ := GetAllInGroup(grid, x, y, c, [][]int{}, [][]int{})

	// copy the grid and remove all cells not in this group
	// surround the group with empty cells (.)
	newGrid := make([][]string, 0)
	for y := 0; y < len(grid)+2; y++ {
		newGrid = append(newGrid, make([]string, 0))
		for x := 0; x < len(grid[0])+2; x++ {
			newGrid[y] = append(newGrid[y], ".")
		}
	}

	for _, c := range cellsInGroup {
		newGrid[c[1]+1][c[0]+1] = grid[c[1]][c[0]]
	}

	grid = newGrid

	for p := 0; p < len(grid)-1; p++ {
		for q := 0; q < len(grid[p])-1; q++ {
			topLeft := grid[p][q]
			topRight := grid[p][q+1]
			bottomLeft := grid[p+1][q]
			bottomRight := grid[p+1][q+1]

			lijstje := []string{topLeft, topRight, bottomLeft, bottomRight}
			hoeveelheid := len(slices.DeleteFunc(lijstje, func(s string) bool {
				return len(s) == 0 || s != c
			}))

			if hoeveelheid == 3 || hoeveelheid == 1 {
				corners++
			}

			if hoeveelheid == 2 && topLeft == bottomRight {
				corners += 2
			}
		}
	}

	return corners
}

func GetEdgeCells(grid [][]string, x, y int) [][]int {
	edges := make([][]int, 0)
	c := grid[y][x]

	group, _ := GetAllInGroup(grid, x, y, c, [][]int{}, [][]int{})
	for _, g := range group {
		neighbors := GetNumberOfNeighbors(grid, g[0], g[1])
		if neighbors < 4 {
			edges = append(edges, g)
		}
	}

	return edges
}

func GetAllInGroup(grid [][]string, x, y int, c string, visited [][]int, group [][]int) ([][]int, [][]int) {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return group, visited
	}

	if grid[y][x] != c {
		return group, visited
	}

	for i := 0; i < len(visited); i += 1 {
		if visited[i][0] == x && visited[i][1] == y {
			return group, visited
		}
	}

	visited = append(visited, []int{x, y})
	group = append(group, []int{x, y})

	if x > 0 && grid[y][x-1] == c {
		group, visited = GetAllInGroup(grid, x-1, y, c, visited, group)
	}
	if x < len(grid[0])-1 && grid[y][x+1] == c {
		group, visited = GetAllInGroup(grid, x+1, y, c, visited, group)
	}
	if y > 0 && grid[y-1][x] == c {
		group, visited = GetAllInGroup(grid, x, y-1, c, visited, group)
	}
	if y < len(grid)-1 && grid[y+1][x] == c {
		group, visited = GetAllInGroup(grid, x, y+1, c, visited, group)
	}

	return group, visited
}

func GetNumberOfNeighbors(grid [][]string, x, y int) int {
	neighbors := 0
	if x > 0 && grid[y][x-1] != "." {
		neighbors++
	}
	if x < len(grid[y])-1 && grid[y][x+1] != "." {
		neighbors++
	}
	if y > 0 && grid[y-1][x] != "." {
		neighbors++
	}
	if y < len(grid)-1 && grid[y+1][x] != "." {
		neighbors++
	}
	return neighbors
}
