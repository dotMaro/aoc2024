package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/dotMaro/aoc2024/utils"
)

func main() {
	listA, listB := parseLocationIDs(utils.SplitInput("day01/input.txt"))
	distances := distances(listA, listB)
	var sum int
	for _, distance := range distances {
		sum += distance
	}
	fmt.Println("Part 1.", sum)
	totalSimilarityScore := 0
	for _, item := range listA {
		totalSimilarityScore += similarityScore(item, listB)
	}
	fmt.Println("Part 2.", totalSimilarityScore)
}

func parseLocationIDs(lines []string) ([]int, []int) {
	listA, listB := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		itemA, itemB, couldSplit := strings.Cut(line, "   ")
		if !couldSplit {
			panic("could not split line " + line)
		}
		var err error
		listA[i], err = strconv.Atoi(itemA)
		if err != nil {
			panic(itemA + " is not a number")
		}
		listB[i], err = strconv.Atoi(itemB)
		if err != nil {
			panic(itemB + " is not a number")
		}
	}
	return listA, listB
}

func distances(listA []int, listB []int) []int {
	slices.Sort(listA)
	slices.Sort(listB)
	distances := make([]int, len(listA))
	for i, itemA := range listA {
		itemB := listB[i]
		distance := itemA - itemB
		if distance < 0 {
			distance *= -1
		}
		distances[i] = distance
	}
	return distances
}

func similarityScore(number int, list []int) int {
	matches := 0
	for _, item := range list {
		if item == number {
			matches++
		}
	}
	return matches * number
}
