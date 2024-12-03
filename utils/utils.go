package utils

import (
	"bufio"
	"os"
)

func ReadInputFromFile(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	var res []string

	if err != nil {
		return res, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil

}
