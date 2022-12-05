package main

import (
	"aoc/tools"
	"fmt"
	"strings"
)

func firstTask(fileContent []string) (int, int) {
	endScore := 0
	secondScore := 0
	for _, element := range fileContent {
		options := strings.Fields(element)
		switch options[1] {
		case "X":
			endScore += 1
			switch options[0] {
			case "A":
				endScore += 3
				secondScore += 3
			case "B":
				secondScore += 1
			case "C":
				endScore += 6
				secondScore += 2
			}
		case "Y":
			endScore += 2
			secondScore += 3
			switch options[0] {
			case "C":
				secondScore += 3
			case "B":
				endScore += 3
				secondScore += 2
			case "A":
				endScore += 6
				secondScore += 1
			}
		case "Z":
			endScore += 3
			secondScore += 6
			switch options[0] {
			case "A":
				secondScore += 2
			case "C":
				endScore += 3
				secondScore += 1
			case "B":
				endScore += 6
				secondScore += 3
			}
		}
	}
	return endScore, secondScore
}

func main() {
	fileContent := tools.ReadFile("./02/02.txt")
	// fileContent := tools.ReadFile("./02/test.txt")
	fmt.Print(firstTask(fileContent))
}
