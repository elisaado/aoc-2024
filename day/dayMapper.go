package day

type Day struct {
	Number int                 // Day number
	Part1  func(string) string // Part1 function
	Part2  func(string) string // Part2 function
}

var dayMap = map[int]Day{}

func RegisterDay(d Day) {
	dayMap[d.Number] = d
}

func GetDay(day int) Day {
	return dayMap[day]
}
