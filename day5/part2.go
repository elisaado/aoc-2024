package day5

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
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
			pages[firstPage] = Page{Number: firstPage, Dependencies: make([]Page, 0)}
		}

		if _, ok := pages[secondPage]; !ok {
			dependencyArray := append(make([]Page, 0), pages[firstPage])
			pages[secondPage] = Page{Number: secondPage, Dependencies: dependencyArray}
		} else {
			pages[secondPage] = Page{Number: secondPage, Dependencies: append(pages[secondPage].Dependencies, pages[firstPage])}
		}
	}

	updates := strings.Split(splat[1], "\n")

	incorrectUpdates := make([][]Page, 0)
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

		if !correct {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	correctUpdates := make([][]Page, 0)
	for _, update := range incorrectUpdates {
		newUpdate := make([]Page, 0)

		for _, page := range update {
			relevantDependencies := make([]Page, 0)

			for _, dependency := range page.Dependencies {
				if UpdateContainsPage(update, dependency) {
					relevantDependencies = append(relevantDependencies, dependency)
				}
			}

			newUpdate = walkPage(page.Number, pages, newUpdate, relevantDependencies)
		}

		if !checkCorrect(newUpdate) {
			panic("incorrect update after correcting")
		}

		correctUpdates = append(correctUpdates, newUpdate)
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

func UpdateContainsPage(update []Page, page Page) bool {
	for _, p := range update {
		if p.Number == page.Number {
			return true
		}
	}

	return false
}

func walkPage(pageNumber int, pages map[int]Page, update []Page, relevantDependencies []Page) []Page {
	// the pages in page.dependency are not the same as the pages in pages,
	// as they were possibly created without their own dependencies
	// this can be fixed using references (pointers) as dependencies,
	// but this works fine for now
	page := pages[pageNumber]

	if UpdateContainsPage(update, page) {
		return update
	}

	for _, dependency := range page.Dependencies {
		// check if relevant
		if !UpdateContainsPage(relevantDependencies, dependency) {
			continue
		}

		// check if dependency is in update already
		if UpdateContainsPage(update, dependency) {
			continue
		}

		update = walkPage(dependency.Number, pages, update, relevantDependencies)
	}

	update = append(update, page)
	return update
}

func SimplifyUpdate(update []Page) []int {
	simplfiied := make([]int, 0)

	for _, page := range update {
		simplfiied = append(simplfiied, page.Number)
	}

	return simplfiied
}
