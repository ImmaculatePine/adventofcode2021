package day1

import (
	"fmt"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

func Task1() error {
	input, err := utils.ReadInputInts("./day1/input.txt")
	if err != nil {
		return err
	}

	result := 0
	last := input[0]

	for _, v := range input {
		if v > last {
			result++
		}
		last = v
	}

	fmt.Println(result)
	return nil
}

func Task2() error {
	input, err := utils.ReadInputInts("./day1/input.txt")
	if err != nil {
		return err
	}

	var window1 []int
	var window2 []int
	result := 0

	for i, v := range input {
		if len(window1) < 3 {
			window1 = append(window1, v)
		}

		// Skip the first value for the second window
		if i > 0 {
			window2 = append(window2, v)
		}

		if len(window1) == 3 && len(window2) == 3 {
			if sum(window2) > sum(window1) {
				result++
			}
			window1 = window2
			window2 = []int{
				window2[1],
				window2[2],
			}
		}
	}

	fmt.Println(result)
	return nil
}

func sum(arr []int) (res int) {
	for _, v := range arr {
		res += v
	}
	return
}
