package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Entry struct {
	signals []string
	outputs []string
}

type Digit = map[rune]bool

func Task1() error {
	entries, err := readInput()
	if err != nil {
		return nil
	}

	sum := 0
	for _, entry := range entries {
		for _, digit := range entry.outputs {
			ln := len(digit)
			if ln == 2 || ln == 4 || ln == 3 || ln == 7 {
				sum++
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func Task2() error {
	entries, err := readInput()
	if err != nil {
		return nil
	}

	sum := 0
	for _, entry := range entries {
		decodedSignals, err := decode(entry.signals)
		if err != nil {
			return err
		}

		values := make(map[string]int)
		for value, signal := range decodedSignals {
			values[digitToStr(signal)] = value
		}

		str := ""
		for _, output := range entry.outputs {
			sortedOutput := digitToStr(strToDigit(output))
			value, ok := values[sortedOutput]
			if !ok {
				return fmt.Errorf("unknown output %s (original %s)", sortedOutput, output)
			}
			str = str + strconv.Itoa(value)
		}
		val, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		sum += val
	}

	fmt.Println(sum)

	return nil
}

func decode(signals []string) ([]Digit, error) {
	one, err := findBySize(signals, 2)
	if err != nil {
		return nil, err
	}
	four, err := findBySize(signals, 4)
	if err != nil {
		return nil, err
	}
	seven, err := findBySize(signals, 3)
	if err != nil {
		return nil, err
	}
	eight, err := findBySize(signals, 7)
	if err != nil {
		return nil, err
	}

	cfMatches := intersection(one, four)
	if len(cfMatches) != 2 {
		return nil, fmt.Errorf("failed to detect potential c- and f-segments")
	}

	twoOrThreeOrFive := findAllBySize(signals, 5)
	var three Digit
	var twoOrFive []Digit
	for _, maybeThree := range twoOrThreeOrFive {
		if doesContainAll(maybeThree, one) {
			three = maybeThree
		} else {
			twoOrFive = append(twoOrFive, maybeThree)
		}
	}
	if three == nil {
		return nil, fmt.Errorf("failed to detect digit 3")
	}

	adg := intersection(twoOrFive[0], twoOrFive[1])
	if len(adg) != 3 {
		return nil, fmt.Errorf("failed to detect potential a-, d- and g-segments")
	}

	var dMatches []rune
	for _, r := range adg {
		_, ok := four[r]
		if ok {
			dMatches = append(dMatches, r)
		}
	}
	if len(dMatches) != 1 {
		return nil, fmt.Errorf("failed to detect d-segment")
	}

	d := dMatches[0]

	zeroOrSixOrNine := findAllBySize(signals, 6)
	var zero Digit
	var sixOrNine []Digit
	for _, maybeZero := range zeroOrSixOrNine {
		_, ok := maybeZero[d]
		if !ok {
			zero = maybeZero
		} else {
			sixOrNine = append(sixOrNine, maybeZero)
		}
	}
	if zero == nil {
		return nil, fmt.Errorf("failed to detect digit 0")
	}

	var six Digit
	var nine Digit
	for _, digit := range sixOrNine {
		if doesContainAll(digit, one) {
			nine = digit
		} else {
			six = digit
		}
	}
	if six == nil {
		return nil, fmt.Errorf("failed to detect digit 6")
	}
	if nine == nil {
		return nil, fmt.Errorf("failed to detect digit 9")
	}

	var c rune
	for _, r := range cfMatches {
		_, ok := six[r]
		if !ok {
			c = r
		}
	}

	var two Digit
	var five Digit
	for _, digit := range twoOrFive {
		_, ok := digit[c]
		if ok {
			two = digit
		} else {
			five = digit
		}
	}

	return []Digit{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}, nil
}

func doesContainAll(superset Digit, subset Digit) bool {
	for r := range subset {
		_, ok := superset[r]
		if !ok {
			return false
		}
	}
	return true
}

func intersection(shorter Digit, longer Digit) []rune {
	res := []rune{}
	for k, _ := range longer {
		_, ok := shorter[k]
		if ok {
			res = append(res, k)
		}
	}
	return res
}

func findAllBySize(signals []string, size int) []Digit {
	var digits []Digit
	for _, signal := range signals {
		if len(signal) == size {
			digits = append(digits, strToDigit(signal))
		}
	}
	return digits
}

func findBySize(signals []string, size int) (Digit, error) {
	for _, signal := range signals {
		if len(signal) == size {
			return strToDigit(signal), nil
		}
	}
	return nil, fmt.Errorf("failed to find signal of size %d", size)
}

func strToDigit(str string) Digit {
	digit := make(Digit)
	for _, r := range str {
		digit[r] = true
	}
	return digit
}

func digitToStr(digit Digit) string {
	var segments []string
	for r := range digit {
		segments = append(segments, string(r))
	}
	sort.Strings(segments)
	return strings.Join(segments, "")
}

func readInput() ([]Entry, error) {
	lines, err := utils.ReadInputStrings("./day8/input.txt")
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		signals := strings.Fields(parts[0])
		outputs := strings.Fields(parts[1])
		entries = append(entries, Entry{signals, outputs})
	}
	return entries, nil
}
