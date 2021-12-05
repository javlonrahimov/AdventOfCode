package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileLineByLine(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		lines = append(lines, data)
	}

	return lines
}

func ReadFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		lines = append(lines, data)
	}

	return strings.Join(lines, "\n")
}
