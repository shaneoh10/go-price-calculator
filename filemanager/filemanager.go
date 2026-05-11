package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read lines from file.")
	}

	file.Close()

	return lines, nil
}
