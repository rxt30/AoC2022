package main

import (
	"aoc/tools"
	"fmt"
)

func readGridContent(fileContent []string) [][]int {
	var treeGrid [][]int = make([][]int, len(fileContent))
	for index, element := range fileContent {
		fmt.Println(element)
		for subIndex, subElement := range element {
			fmt.Println(index)
			fmt.Println(subIndex)
			fmt.Println(int(subElement))
			treeGrid[index][subIndex] = int(subElement)
		}
	}
	return treeGrid
}

func firstTask(fileContent []string) int {
	treeGrid := readGridContent(fileContent)
	fmt.Print(treeGrid)
	fmt.Print(treeGrid[1:][1])
	var visibleTress int = len(treeGrid)*2 + len(treeGrid[0])*2 - 4
	return visibleTress
}

func main() {
	fileContent := tools.ReadFile("./08/test.txt")
	// fileContent := tools.ReadFile("./08/08.txt")
	fmt.Print(firstTask(fileContent))
}
