package day8

import (
	"strconv"
	"strings"
)

type Point struct {
	x        int
	y        int
	antenna  string
	antinode bool
}

func (p Point) String() string {
	return "Point{x: " + strconv.Itoa(p.x) + ", y: " + strconv.Itoa(p.y) + ", antenna: " + p.antenna + ", antinode: " + strconv.FormatBool(p.antinode) + "}"
}

type Grid [][]Point

func (g Grid) String() string {
	str := ""
	for _, row := range g {
		for _, point := range row {
			chr := point.antenna
			if point.antinode {
				chr = "#"
			}
			str += chr
		}
		str += "\n"
	}
	return str
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	grid := make(Grid, 0)
	antennas := make([]Point, 0)
	for y, line := range lines {
		if line == "" {
			continue
		}
		chars := strings.Split(line, "")
		row := make([]Point, 0)
		for x, char := range chars {
			point := Point{x, y, char, false}
			row = append(row, point)
			if char != "." && char != "#" {
				antennas = append(antennas, point)
			}
		}
		grid = append(grid, row)
	}

	for _, antenna := range antennas {
		println(antenna.String())
	}
	println()

	donePairs := make(map[string]bool)
	for _, antenna := range antennas {
		for _, otherAntenna := range antennas {
			if antenna == otherAntenna {
				continue
			}

			if antenna.antenna == otherAntenna.antenna {
				if donePairs[HashPair(antenna, otherAntenna)] {
					continue
				}
				xDiff := Abs(antenna.x - otherAntenna.x)
				yDiff := Abs(antenna.y - otherAntenna.y)

				antinodeOne := Point{}
				antinodeTwo := Point{}

				if antenna.x < otherAntenna.x {
					xDiff = -xDiff
				}

				if antenna.y < otherAntenna.y {
					yDiff = -yDiff
				}

				// if antenna.x < otherAntenna.x && antenna.y < otherAntenna.y {
				// 	antinodeOne = Point{antenna.x - xDiff, antenna.y - yDiff, "#", true}
				// 	antinodeTwo = Point{otherAntenna.x + xDiff, otherAntenna.y + yDiff, "#", true}
				// }

				// if antenna.x < otherAntenna.x && antenna.y > otherAntenna.y {
				// 	antinodeOne = Point{antenna.x - xDiff, antenna.y + yDiff, "#", true}
				// 	antinodeTwo = Point{otherAntenna.x + xDiff, otherAntenna.y - yDiff, "#", true}
				// }

				// if antenna.x > otherAntenna.x && antenna.y < otherAntenna.y {
				// 	antinodeOne = Point{antenna.x + xDiff, antenna.y - yDiff, "#", true}
				// 	antinodeTwo = Point{otherAntenna.x - xDiff, otherAntenna.y + yDiff, "#", true}
				// }

				// if antenna.x > otherAntenna.x && antenna.y > otherAntenna.y {
				// 	antinodeOne = Point{antenna.x + xDiff, antenna.y + yDiff, "#", true}
				// 	antinodeTwo = Point{otherAntenna.x - xDiff, otherAntenna.y - yDiff, "#", true}
				// }

				antinodeOne = Point{antenna.x + xDiff, antenna.y + yDiff, "#", true}
				antinodeTwo = Point{otherAntenna.x - xDiff, otherAntenna.y - yDiff, "#", true}

				if antinodeOne.x < len(grid[0]) && antinodeOne.y < len(grid) && antinodeOne.x >= 0 && antinodeOne.y >= 0 {
					antinodeOne.antenna = grid[antinodeOne.y][antinodeOne.x].antenna
					grid[antinodeOne.y][antinodeOne.x] = antinodeOne
				}

				if antinodeTwo.x < len(grid[0]) && antinodeTwo.y < len(grid) && antinodeTwo.x >= 0 && antinodeTwo.y >= 0 {
					antinodeTwo.antenna = grid[antinodeTwo.y][antinodeTwo.x].antenna
					grid[antinodeTwo.y][antinodeTwo.x] = antinodeTwo
				}

				donePairs[HashPair(antenna, otherAntenna)] = true
			}
		}
	}

	println(grid.String())

	count := 0
	for _, row := range grid {
		for _, point := range row {
			if point.antinode {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func HashPair(a, b Point) string {
	if a.x > b.x {
		return a.String() + b.String()
	}
	return b.String() + a.String()
}
