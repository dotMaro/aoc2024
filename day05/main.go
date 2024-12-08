package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/aoc2024/utils"
)

func main() {
	lines := utils.SplitInput("day05/input.txt")
	manualUpdates := parseManualUpdates(lines)
	correctUpdates, incorrectUpdates := manualUpdates.findCorrectlyAndIncorrectlyOrderedUpdates()
	sumOfCorrectMiddlePages := 0
	for _, u := range correctUpdates {
		middlePage := u[len(u)/2]
		sumOfCorrectMiddlePages += middlePage
	}
	fmt.Println("Part 1.", sumOfCorrectMiddlePages)
	reorderedUpdates := manualUpdates.reorderedUpdates(incorrectUpdates)
	sumOfCorrectedMiddlePages := 0
	for _, u := range reorderedUpdates {
		middlePage := u[len(u)/2]
		sumOfCorrectedMiddlePages += middlePage
	}
	fmt.Println("Part 2.", sumOfCorrectedMiddlePages)
}

func parseManualUpdates(lines []string) manualUpdates {
	rulesPart := true
	pageOrderingRules := make(map[int][]int, len(lines))
	updates := make([][]int, 0, 70)
	for _, line := range lines {
		if line == "" {
			rulesPart = false
			continue
		}
		if rulesPart {
			left, err := strconv.Atoi(line[:2])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(line[3:])
			if err != nil {
				panic(err)
			}
			pageOrderingRules[right] = append(pageOrderingRules[right], left)
		} else {
			values := strings.Split(line, ",")
			update := make([]int, len(values))
			for i, value := range values {
				var err error
				update[i], err = strconv.Atoi(value)
				if err != nil {
					panic(err)
				}
			}
			updates = append(updates, update)
		}
	}

	return manualUpdates{
		pageOrderingRules: pageOrderingRules,
		updates:           updates,
	}
}

type manualUpdates struct {
	pageOrderingRules map[int][]int // right -> left
	updates           [][]int
}

func (u manualUpdates) findCorrectlyAndIncorrectlyOrderedUpdates() (correct [][]int, incorrect [][]int) {
	correctUpdates := make([][]int, 0, len(u.updates))
	incorrectUpdates := make([][]int, 0, len(u.updates))
	for _, update := range u.updates {
		hasPrinted := make(map[int]bool, len(update)) // true if printed, false if they will be printed
		for _, page := range update {
			hasPrinted[page] = false
		}

		correct := true
		for _, page := range update {
			shouldBePrintedList := u.pageOrderingRules[page]
			for _, shouldBePrinted := range shouldBePrintedList {
				hasBeenPrinted, includedInUpdate := hasPrinted[shouldBePrinted]
				if includedInUpdate && !hasBeenPrinted {
					correct = false
					break
				}
			}
			hasPrinted[page] = true
		}
		if correct {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	return correctUpdates, incorrectUpdates
}

func (u manualUpdates) reorderedUpdates(updates [][]int) [][]int {
	reorderedUpdates := make([][]int, len(updates))
	for i, update := range updates {
		reorderedUpdates[i] = u.reorderUpdateToBeCorrect(update)
	}
	return reorderedUpdates
}

func (u manualUpdates) reorderUpdateToBeCorrect(update []int) []int {
	hasPrinted := make(map[int]bool, len(update)) // true if printed, false if they will be printed
	for _, page := range update {
		hasPrinted[page] = false
	}

	reordered := make([]int, 0, len(update))
	for len(reordered) < len(update) {
		for _, page := range update {
			if hasPrinted[page] {
				continue
			}
			noUnfulfilledRules := true
			shouldBePrintedList := u.pageOrderingRules[page]
			for _, shouldBePrinted := range shouldBePrintedList {
				hasBeenPrinted, includedInUpdate := hasPrinted[shouldBePrinted]
				if includedInUpdate && !hasBeenPrinted {
					noUnfulfilledRules = false
					break
				}
			}
			if noUnfulfilledRules {
				hasPrinted[page] = true
				reordered = append(reordered, page)
				break
			}
		}
	}

	return reordered
}
