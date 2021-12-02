package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileLineByLine(filePath string) []int64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, int64(data))
	}

	return lines
}