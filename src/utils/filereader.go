package utils

import (
	"bufio"
	"errors"
	"os"
)

func GetLinesFromTextFile(input string) ([]string, error) {
	file, err := os.Open(input)
	defer file.Close()

	if err != nil {
		return nil, errors.New("file not found")
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}
