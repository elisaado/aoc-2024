package day5

import (
	"strconv"
	"strings"
)

type Page struct {
	Number       int
	Dependencies []Page
}

func Part1(input string) string {
	splat := strings.Split(input, "\n\n")
	order := splat[0]

	pages := make(map[int]Page)

	for _, line := range strings.Split(order, "\n") {
		if line == "" {
			continue
		}

		pagesInLine := strings.Split(line, "|")
		firstPageStr := strings.TrimSpace(pagesInLine[0])
		firstPage, err := strconv.Atoi(firstPageStr)
		if err != nil {
			panic(err)
		}

		secondPageStr := strings.TrimSpace(pagesInLine[1])
		secondPage, err := strconv.Atoi(secondPageStr)
		if err != nil {
			panic(err)
		}

		if _, ok := pages[firstPage]; !ok {
			pages[firstPage] = Page{Number: firstPage}
		}

		if _, ok := pages[secondPage]; !ok {
			pages[secondPage] = Page{Number: secondPage, Dependencies: []Page{pages[firstPage]}}
		} else {
			pages[secondPage] = Page{Number: secondPage, Dependencies: append(pages[secondPage].Dependencies, pages[firstPage])}
		}
	}

	updates := strings.Split(splat[1], "\n")

	correctUpdates := make([][]Page, 0)
	for _, updateLine := range updates {
		if updateLine == "" {
			continue
		}

		updateStrs := strings.Split(updateLine, ",")
		update := make([]Page, 0)

		for _, updateStr := range updateStrs {
			updateStr = strings.TrimSpace(updateStr)
			updateInt, err := strconv.Atoi(updateStr)
			if err != nil {
				panic(err)
			}

			update = append(update, pages[updateInt])
		}

		correct := checkCorrect(update)

		if correct {
			correctUpdates = append(correctUpdates, update)
		}
	}

	sum := 0
	for _, update := range correctUpdates {
		if len(update)%2 == 0 {
			panic("even number of updates")
		}

		middle := update[len(update)/2]
		sum += middle.Number
	}

	return strconv.Itoa(sum)
}

func checkCorrect(update []Page) bool {
	for i, page := range update {
		for _, dep := range page.Dependencies {
			for j := i; j < len(update); j++ {
				if update[j].Number == dep.Number {
					return false
				}
			}
		}
	}

	return true
}
