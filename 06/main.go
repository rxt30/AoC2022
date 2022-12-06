package main

import (
	"aoc/tools"
	"fmt"
)

func uniqueCharactersInString(stringToCheck string) bool {
	charDict := make(map[rune]int)
	for _, character := range stringToCheck {
		charDict[character] = 1
	}
	return len(charDict) == len(stringToCheck)
}

func findTask(fileContent []string, lengtOfString int) int {
	packageString := fileContent[0]
	for i := 0; i < len(packageString)-lengtOfString; i++ {
		if uniqueCharactersInString(packageString[i : i+lengtOfString]) {
			return i + lengtOfString
		}
	}
	return 0
}

func main() {
	// fileContent := tools.ReadFile("./06/test.txt")
	fileContent := tools.ReadFile("./06/06.txt")
	fmt.Println(findTask(fileContent, 4))
	fmt.Println(findTask(fileContent, 14))
}
