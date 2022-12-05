package main

import (
	"aoc/tools"
	"fmt"
	"sort"
	"strconv"
)

func firstProblem(fileContent []string) (int, int) {
	var calories []int
	var currentElement int
	currentCalories := 0
	for _, element := range fileContent {
		if len(element) == 0 {
			calories = append(calories, currentCalories)
			currentCalories = 0
		} else {
			currentElement, _ = strconv.Atoi(element)
			currentCalories += currentElement
		}
	}
	sort.Ints(calories)
	return calories[len(calories)-1], (calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])
}

func main() {
	fileContent := tools.ReadFile("./01/01.txt")
	fmt.Print(firstProblem(fileContent))
}
