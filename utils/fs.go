package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ReadInputAsCommaSeparatedInts(path string) ([]int, error) {
	input, err := ReadInputStrings(path)
	if err != nil {
		return nil, err
	}

	if len(input) != 1 {
		return nil, fmt.Errorf("unexpected number of input lines %d", len(input))
	}

	parts := strings.Split(input[0], ",")
	result := []int{}

	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}

	return result, nil
}
