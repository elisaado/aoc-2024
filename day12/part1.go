package day12

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
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
			alreadyVisited := false
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
			circum := GetCircum(grid, x, y)
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
			println("Letter area:", vv[0], "Letter circum:", vv[1])
			sum += vv[0] * vv[1]
		}
	}

	return strconv.Itoa(sum)
}

func GetArea(grid [][]string, x int, y int) (int, [][]int) {
	c := grid[y][x]

	visited := make([][]int, 0)
	area, visited := GetAreaRecursive(grid, c, x, y, visited)

	return area, visited
}

func GetAreaRecursive(grid [][]string, c string, x int, y int, visited [][]int) (int, [][]int) {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0, visited
	}

	for _, v := range visited {
		if v[0] == x && v[1] == y {
			return 0, visited
		}
	}

	if grid[y][x] != c {
		return 0, visited
	}
	visited = append(visited, []int{x, y})

	// recurse over all neighbors
	left, visited := GetAreaRecursive(grid, c, x-1, y, visited)
	right, visited := GetAreaRecursive(grid, c, x+1, y, visited)
	up, visited := GetAreaRecursive(grid, c, x, y-1, visited)
	down, visited := GetAreaRecursive(grid, c, x, y+1, visited)

	return 1 + left + right + up + down, visited
}

func GetCircum(grid [][]string, x int, y int) int {
	i, _ := GetCircumRecursive(grid, grid[y][x], x, y, [][]int{})
	return i
}

func GetCircumRecursive(grid [][]string, c string, x int, y int, visited [][]int) (int, [][]int) {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0, visited
	}

	for _, v := range visited {
		if v[0] == x && v[1] == y {
			return 0, visited
		}
	}

	if grid[y][x] != c {
		return 0, visited
	}
	visited = append(visited, []int{x, y})

	surfaceArea := 0

	// check left
	if x-1 < 0 || grid[y][x-1] != c {
		surfaceArea++
		println("x, y", x, y)
		println(surfaceArea)
	} else {
		surfaceAreaLeft, visitedLeft := GetCircumRecursive(grid, c, x-1, y, visited)
		visited = visitedLeft
		surfaceArea += surfaceAreaLeft
	}

	// check right
	if x+1 >= len(grid[0]) || grid[y][x+1] != c {
		surfaceArea++
	} else {
		surfaceAreaRight, visitedRight := GetCircumRecursive(grid, c, x+1, y, visited)
		visited = visitedRight
		surfaceArea += surfaceAreaRight
	}

	// check up
	if y-1 < 0 || grid[y-1][x] != c {
		surfaceArea++
	} else {
		surfaceAreaUp, visitedUp := GetCircumRecursive(grid, c, x, y-1, visited)
		visited = visitedUp
		surfaceArea += surfaceAreaUp
	}

	// check down
	if y+1 >= len(grid) || grid[y+1][x] != c {
		surfaceArea++
	} else {
		surfaceAreaDown, visitedDown := GetCircumRecursive(grid, c, x, y+1, visited)
		visited = visitedDown
		surfaceArea += surfaceAreaDown
	}

	return surfaceArea, visited
}
