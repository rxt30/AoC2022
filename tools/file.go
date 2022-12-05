package tools

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) []string {
	var fileContent []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	return fileContent
}
