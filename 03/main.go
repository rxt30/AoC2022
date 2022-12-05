package main

import (
	"aoc/tools"
	"fmt"
	"strings"
	"unicode"
)

func getDuplicateItems(firstHalf string, secondHalf string) map[rune]int {
	foundItems := make(map[rune]int)
	var charValue int
	for _, leftChar := range firstHalf {
		if strings.ContainsRune(secondHalf, leftChar) {
			charValue = int(leftChar) - 96
			if unicode.IsUpper(leftChar) {
				charValue = int(leftChar) - 38
			}
			foundItems[leftChar] = charValue
		}
	}
	return foundItems
}

func getDuplicateItemsForSecondTask(firstHalf string, secondHalf string, thirdHalf string) map[rune]int {
	foundItems := make(map[rune]int)
	var charValue int
	for _, leftChar := range firstHalf {
		if strings.ContainsRune(secondHalf, leftChar) && strings.ContainsRune(thirdHalf, leftChar) {
			charValue = int(leftChar) - 96
			if unicode.IsUpper(leftChar) {
				charValue = int(leftChar) - 38
			}
			foundItems[leftChar] = charValue
		}
	}
	return foundItems
}

func firstTask(fileContent []string) int {
	totalScore := 0
	for _, element := range fileContent {
		firstHalf := element[0 : len(element)/2]
		secondHalf := element[len(element)/2:]
		foundCharacters := getDuplicateItems(firstHalf, secondHalf)
		for _, val := range foundCharacters {
			totalScore += val
		}
	}
	return totalScore
}

func secondTask(fileContent []string) int {
	totalScore := 0
	for i := 0; i < len(fileContent)-1; i += 3 {
		foundCharacters := getDuplicateItemsForSecondTask(fileContent[i], fileContent[i+1], fileContent[i+2])
		for _, val := range foundCharacters {
			totalScore += val
		}
	}
	return totalScore
}

func main() {
	// fileContent := tools.ReadFile("./03/test.txt")
	fileContent := tools.ReadFile("./03/03.txt")
	fmt.Print(firstTask(fileContent), secondTask(fileContent))
}
