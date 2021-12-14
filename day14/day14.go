package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

func Task1() error {
	result, err := solve(10)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func Task2() error {
	result, err := solve(40)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func solve(steps int) (int, error) {
	initial, diff, err := readInput()
	if err != nil {
		return 0, err
	}

	elements := make(map[string]int)
	patterns := make(map[string]int)
	pattern := ""
	for _, char := range initial {
		element := string(char)
		elements[element]++

		if len(pattern) < 2 {
			pattern += element
		} else if len(pattern) == 2 {
			pattern = string(pattern[1]) + element
		}

		if len(pattern) == 2 {
			patterns[pattern]++
		}
	}

	for i := 0; i < steps; i++ {
		apply(elements, patterns, diff)
	}

	min, max, err := findMinMax(elements)
	if err != nil {
		return 0, err
	}

	return max - min, nil
}

func apply(elements map[string]int, patterns map[string]int, diff map[string]string) {
	changes := make(map[string]int)
	for pattern, newEl := range diff {
		if patterns[pattern] > 0 {
			elements[newEl] += patterns[pattern]
			changes[pattern] -= patterns[pattern]
			newPattern1 := string(pattern[0]) + newEl
			newPattern2 := newEl + string(pattern[1])
			changes[newPattern1] += patterns[pattern]
			changes[newPattern2] += patterns[pattern]
		}
	}

	for pattern, val := range changes {
		patterns[pattern] += val
		if patterns[pattern] < 0 {
			patterns[pattern] = 0
		}
	}
}

func findMinMax(counters map[string]int) (int, int, error) {
	vals := make([]int, 0, len(counters))
	for _, v := range counters {
		vals = append(vals, v)
	}

	min, err := utils.FindMinInt(vals)
	if err != nil {
		return 0, 0, err
	}

	max, err := utils.FindMaxInt(vals)
	if err != nil {
		return 0, 0, err
	}

	return min, max, nil
}

func readInput() (string, map[string]string, error) {
	file, err := os.Open("./day14/input.txt")
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan initial pattern
	scanner.Scan()
	initial := scanner.Text()

	// Skip empty line
	scanner.Scan()

	// Scan instructions
	r, err := regexp.Compile("^([A-Z]{2}) -> ([A-Z])$")
	if err != nil {
		return "", nil, err
	}

	diff := make(map[string]string)
	for scanner.Scan() {
		matches := r.FindStringSubmatch(scanner.Text())
		if len(matches) != 3 {
			return "", nil, fmt.Errorf("failed to parse instruction %s", scanner.Text())
		}

		diff[matches[1]] = matches[2]
	}

	return initial, diff, scanner.Err()
}
