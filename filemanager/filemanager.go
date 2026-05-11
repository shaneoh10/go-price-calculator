package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJSON(filePath string, data interface{}) error {
	file, err := os.Create(filePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to write JSON to file.")
	}

	file.Close()
	return nil
}
