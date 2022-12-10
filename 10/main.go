package main

import (
	"aoc/tools"
	"fmt"
	"strconv"
	"strings"
)

func getCycleValues(fileContent []string) []int {
	var cycleValue []int
	cycleValue = append(cycleValue, 1)
	for _, element := range fileContent {
		currentLine := strings.Fields(element)
		switch currentLine[0] {
		case "noop":
			cycleValue = append(cycleValue, cycleValue[len(cycleValue)-1])
		case "addx":
			cycleValue = append(cycleValue, cycleValue[len(cycleValue)-1])
			// cycleValue = append(cycleValue, cycleValue[len(cycleValue)-1])
			incrementValue, _ := strconv.Atoi(currentLine[1])
			cycleValue = append(cycleValue, cycleValue[len(cycleValue)-1]+incrementValue)
		}
	}
	return cycleValue
}

func drawText(cycleValue []int) {
	for i := 1; i <= 240; i++ {
		currentValue := cycleValue[i-1] + 1
		if i%40 >= currentValue-1 && i%40 <= currentValue+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if i%40 == 0 {
			fmt.Print("\n")
		}
	}
}

func getCyclePointReturn(cycleValues []int) int {
	var cycleValuesSum int = 0
	for i := 20; i < len(cycleValues); i += 40 {
		fmt.Println(cycleValues[i-1] * i)
		cycleValuesSum += cycleValues[i-1] * i
	}
	return cycleValuesSum
}

func firstTask(fileContent []string) int {
	cycleValues := getCycleValues(fileContent)
	return getCyclePointReturn(cycleValues)
}

func secondTask(fileContent []string) {
	cycleValues := getCycleValues(fileContent)
	drawText(cycleValues)
}
func main() {
	fileContent := tools.ReadFile("./10/10.txt")
	// fileContent := tools.ReadFile("./10/test.txt")
	fmt.Println(firstTask(fileContent))
	secondTask(fileContent)
}
