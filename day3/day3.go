package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

const size = 12

func Task1() error {
	input, err := utils.ReadInputStrings("./day3/input.txt")
	if err != nil {
		return err
	}

	result, err := task1(input)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func task1(input []string) (int64, error) {
	var frequencies [size]struct {
		zeros uint
		ones  uint
	}

	for _, binary := range input {
		for j := 0; j < size; j++ {
			bit := binary[j]
			if bit == '0' {
				frequencies[j].zeros++
			} else if bit == '1' {
				frequencies[j].ones++
			}
		}
	}

	var gammaBits [size]string
	var epsilonBits [size]string

	for i, freq := range frequencies {
		if freq.zeros > freq.ones {
			gammaBits[i] = "0"
			epsilonBits[i] = "1"
		} else {
			gammaBits[i] = "1"
			epsilonBits[i] = "0"
		}
	}

	gamma, err := strconv.ParseInt(strings.Join(gammaBits[:], ""), 2, 64)
	if err != nil {
		return 0, err
	}

	epsilon, err := strconv.ParseInt(strings.Join(epsilonBits[:], ""), 2, 64)
	if err != nil {
		return 0, err
	}

	return gamma * epsilon, nil
}

func Task2() error {
	input, err := utils.ReadInputStrings("./day3/input.txt")
	if err != nil {
		return err
	}

	result, err := task2(input)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func task2(binaries []string) (int64, error) {
	gamma, err := byMostCommon(binaries, 0)
	if err != nil {
		return 0, err
	}
	epsilon, err := byLeastCommon(binaries, 0)
	if err != nil {
		return 0, err
	}

	return gamma * epsilon, nil
}

func byMostCommon(binaries []string, bitNum int) (int64, error) {
	if len(binaries) == 1 {
		return strconv.ParseInt(binaries[0], 2, 64)
	}

	zeros, ones := group(binaries, bitNum)
	if len(ones) >= len(zeros) {
		return byMostCommon(ones, bitNum+1)
	}
	return byMostCommon(zeros, bitNum+1)
}

func byLeastCommon(binaries []string, bitNum int) (int64, error) {
	if len(binaries) == 1 {
		return strconv.ParseInt(binaries[0], 2, 64)
	}

	zeros, ones := group(binaries, bitNum)
	if len(zeros) <= len(ones) {
		return byLeastCommon(zeros, bitNum+1)
	}
	return byLeastCommon(ones, bitNum+1)
}

func group(binaries []string, bitNum int) (zeros, ones []string) {
	for _, binary := range binaries {
		bit := binary[bitNum]
		if bit == '0' {
			zeros = append(zeros, binary)
		} else if bit == '1' {
			ones = append(ones, binary)
		}
	}
	return
}
