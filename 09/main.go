package main

import (
	"aoc/tools"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func removeDuplicates(positions []Position) map[Position]int {
	var uniquePositions map[Position]int = make(map[Position]int)
	for _, element := range positions {
		uniquePositions[element] = 0
	}
	return uniquePositions
}

func getPositions(fileContent []string) []Position {
	var positions []Position
	positions = append(positions, Position{0, 0})
	for _, element := range fileContent {
		currentMovement := strings.Fields(element)
		direction := currentMovement[0]
		distance, _ := strconv.Atoi(currentMovement[1])
		for i := 0; i < distance; i++ {
			lastPosition := positions[len(positions)-1]
			switch direction {
			case "R":
				positions = append(positions, Position{lastPosition.x + 1, lastPosition.y})
			case "L":
				positions = append(positions, Position{lastPosition.x - 1, lastPosition.y})
			case "U":
				positions = append(positions, Position{lastPosition.x, lastPosition.y + 1})
			case "D":
				positions = append(positions, Position{lastPosition.x, lastPosition.y - 1})
			}
		}
	}
	return positions
}

func firstTask(fileContent []string) int {
	var positions []Position = getPositions(fileContent)
	return len(removeDuplicates(positionsCreating(positions)))
}

func secondTask(fileContent []string) int {
	var positions []Position = getPositions(fileContent)
	for i := 0; i < 9; i++ {
		positions = positionsCreating(positions)
	}
	return len(removeDuplicates(positions))
}

func positionsCreating(positions []Position) []Position {
	var tailPositions []Position
	tailPositions = append(tailPositions, Position{0, 0})
	for _, position := range positions[1:] {
		tailPosition := tailPositions[len(tailPositions)-1]
		xChangeAbs := int(math.Abs(float64(tailPosition.x - position.x)))
		yChangeAbs := int(math.Abs(float64(tailPosition.y - position.y)))
		if xChangeAbs > 1 || yChangeAbs > 1 {
			var newX int
			var newY int
			xChange := tailPosition.x - position.x
			yChange := tailPosition.y - position.y
			switch {
			case xChange == 0:
				newX = tailPosition.x
			case xChange > 0:
				newX = tailPosition.x - 1
			case xChange < 0:
				newX = tailPosition.x + 1
			}
			switch {
			case yChange == 0:
				newY = tailPosition.y
			case yChange > 0:
				newY = tailPosition.y - 1
			case yChange < 0:
				newY = tailPosition.y + 1
			}
			tailPositions = append(tailPositions, Position{newX, newY})
		}
	}
	return tailPositions
}

func main() {
	// fileContent := tools.ReadFile("./09/test.txt")
	fileContent := tools.ReadFile("./09/09.txt")
	fmt.Println(firstTask(fileContent))
	fmt.Println(secondTask(fileContent))
}
