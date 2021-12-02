package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadInputInts(path string) ([]int, error) {
	lines, err := ReadInputStrings(path)
	if err != nil {
		return nil, err
	}

	var vals []int
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}

	return vals, nil
}
