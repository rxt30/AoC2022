package main

import (
	"aoc/tools"
	"fmt"
	"strconv"
	"strings"
)

func firstTask(fileContent []string) (int, int) {
	validPairs := 0
	secondValidPairs := 0
	for _, element := range fileContent {
		currentPairs := strings.Split(element, ",")
		leftRange := strings.Split(currentPairs[0], "-")
		rightRange := strings.Split(currentPairs[1], "-")
		leftStart, _ := strconv.Atoi(leftRange[0])
		leftEnd, _ := strconv.Atoi(leftRange[1])
		rightStart, _ := strconv.Atoi(rightRange[0])
		rightEnd, _ := strconv.Atoi(rightRange[1])
		leftIncluded := (leftStart >= rightStart) && (leftEnd <= rightEnd)
		rightIncluded := (leftStart <= rightStart) && (leftEnd >= rightEnd)
		// endComparison := leftEnd <= rightEnd
		leftSecond := leftStart <= rightEnd
		rightSecond := rightStart <= leftEnd
		if leftIncluded || rightIncluded {
			validPairs++
		}
		if leftSecond && rightSecond {
			secondValidPairs++
		}
	}
	return validPairs, secondValidPairs
}

func main() {
	fileContent := tools.ReadFile("./04/04.txt")
	// fileContent := tools.ReadFile("./04/test.txt")
	fmt.Print(firstTask(fileContent))
}
