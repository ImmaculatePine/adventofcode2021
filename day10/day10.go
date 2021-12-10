package day10

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type ErrCorrupted struct {
	Got string
}

func (err ErrCorrupted) Error() string {
	return fmt.Sprintf("line is corrupted, got %s", err.Got)
}

var openingChars = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
	"<": true,
}

var endingToOpeningChars = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var openingToEndingChars = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func Task1() error {
	lines, err := readInput()
	if err != nil {
		return err
	}

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	total := 0
	for _, line := range lines {
		_, err := autocomplete(line)
		if err != nil {
			err, ok := err.(ErrCorrupted)
			if ok {
				total += points[err.Got]
				continue
			}
			return err
		}
	}

	fmt.Println(total)

	return nil
}

func Task2() error {
	lines, err := readInput()
	if err != nil {
		return err
	}

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	scores := []int{}
	for _, line := range lines {
		sequence, err := autocomplete(line)
		if err != nil {
			// Skip corrupted lines
			_, ok := err.(ErrCorrupted)
			if ok {
				continue
			}
			return err
		}

		score := 0
		for _, char := range sequence {
			score *= 5
			score += points[char]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	idx := (len(scores) - 1) / 2
	fmt.Println(scores[idx])

	return nil
}

func autocomplete(line []string) ([]string, error) {
	stack := []string{}
	for _, char := range line {
		if openingChars[char] {
			stack = append(stack, char)
		} else {
			openingChar, ok := endingToOpeningChars[char]
			if !ok {
				return nil, fmt.Errorf("unexpected character %s", char)
			}
			if stack[len(stack)-1] == openingChar {
				stack = stack[:len(stack)-1]
			} else {
				return nil, ErrCorrupted{Got: char}
			}
		}
	}

	autocompleteStack := []string{}
	for i := len(stack) - 1; i >= 0; i-- {
		openingChar := stack[i]
		endingChar, ok := openingToEndingChars[openingChar]
		if !ok {
			return nil, fmt.Errorf("unexpected character %s", openingChar)
		}
		autocompleteStack = append(autocompleteStack, endingChar)
	}
	return autocompleteStack, nil
}

func readInput() ([][]string, error) {
	lines, err := utils.ReadInputStrings("./day10/input.txt")
	if err != nil {
		return nil, err
	}

	res := make([][]string, 0, len(lines))
	for _, line := range lines {
		chars := strings.Split(line, "")
		res = append(res, chars)
	}
	return res, nil
}
