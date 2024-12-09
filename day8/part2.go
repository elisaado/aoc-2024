package day8

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
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

				// antinodes := make([]Point, 0)
				if antenna.x < otherAntenna.x {
					xDiff = -xDiff
				}

				if antenna.y < otherAntenna.y {
					yDiff = -yDiff
				}

				// from antenna
				i := 0
				for {
					x := antenna.x + xDiff*i
					y := antenna.y + yDiff*i
					if x >= len(grid[0]) || y >= len(grid) || x < 0 || y < 0 {
						break
					}

					if otherAntenna.x > antenna.x && x > otherAntenna.x {
						break
					}

					if otherAntenna.x < antenna.x && x < otherAntenna.x {
						break
					}

					if otherAntenna.y > antenna.y && y > otherAntenna.y {
						break
					}

					if otherAntenna.y < antenna.y && y < otherAntenna.y {
						break
					}

					antinode := Point{x, y, "#", true}
					antinode.antenna = grid[antinode.y][antinode.x].antenna
					grid[antinode.y][antinode.x] = antinode

					i++
				}

				// from otherAntenna
				i = 0
				for {
					x := otherAntenna.x - xDiff*i
					y := otherAntenna.y - yDiff*i
					if x >= len(grid[0]) || y >= len(grid) || x < 0 || y < 0 {
						break
					}

					if antenna.x > otherAntenna.x && x > antenna.x {
						break
					}

					if antenna.x < otherAntenna.x && x < antenna.x {
						break
					}

					if antenna.y > otherAntenna.y && y > antenna.y {
						break
					}

					if antenna.y < otherAntenna.y && y < antenna.y {
						break
					}

					antinode := Point{x, y, "#", true}
					antinode.antenna = grid[antinode.y][antinode.x].antenna
					grid[antinode.y][antinode.x] = antinode

					i++
				}

				// antinodeOne := Point{antenna.x + xDiff, antenna.y + yDiff, "#", true}
				// antinodeTwo := Point{otherAntenna.x - xDiff, otherAntenna.y - yDiff, "#", true}

				// if antinodeOne.x < len(grid[0]) && antinodeOne.y < len(grid) && antinodeOne.x >= 0 && antinodeOne.y >= 0 {
				// 	antinodeOne.antenna = grid[antinodeOne.y][antinodeOne.x].antenna
				// 	grid[antinodeOne.y][antinodeOne.x] = antinodeOne
				// }

				// if antinodeTwo.x < len(grid[0]) && antinodeTwo.y < len(grid) && antinodeTwo.x >= 0 && antinodeTwo.y >= 0 {
				// 	antinodeTwo.antenna = grid[antinodeTwo.y][antinodeTwo.x].antenna
				// 	grid[antinodeTwo.y][antinodeTwo.x] = antinodeTwo
				// }

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
