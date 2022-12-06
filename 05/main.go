package main

import (
	"aoc/tools"
	"fmt"
	"strings"
)

func firstTask(fileContent []string) string {
	var stacks [5][5]string
	i := 0
	for ; len(fileContent[i]) != 0; i++ {
		fmt.Print(i)
		currentIndex := 0
		for currentIndex != - {
			currentIndex = strings.IndexRune(fileContent[i][currentIndex:], '[')
			if currentIndex == -1 {
				break
			}
			currentStack := fileContent[i][currentIndex : currentIndex+3]
			fmt.Print(currentIndex)
			fmt.Print(currentStack)
			stacks[currentIndex/4] = append(stacks[currentIndex/4], currentStack)
			fileContent[i] = strings.Replace(fileContent[i], "[", "]", 1)
			if currentIndex != -1 {
				currentIndex++
			}
		}
	}
	fmt.Print(stacks)
	return "Test"
}

func main() {
	fileContent := tools.ReadFile("./05/test.txt")
	// for _, element := range fileContent {
	// 	fmt.Println(len(element))
	// }
	fmt.Print(firstTask(fileContent))
}
