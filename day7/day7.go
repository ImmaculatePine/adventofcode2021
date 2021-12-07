package day7

import (
	"fmt"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

func Task1() error {
	minDiff, err := solveWith(sumOfDiffs)
	if err != nil {
		return err
	}

	fmt.Println(minDiff)

	return nil
}

func Task2() error {
	minDiff, err := solveWith(progressiveSumOfDiffs)
	if err != nil {
		return err
	}

	fmt.Println(minDiff)

	return nil
}

func solveWith(function func([]int, int) int) (int, error) {
	initialVals, err := utils.ReadInputAsCommaSeparatedInts("./day7/input.txt")
	if err != nil {
		return 0, err
	}

	min, err := utils.FindMinInt(initialVals)
	if err != nil {
		return 0, err
	}

	max, err := utils.FindMaxInt(initialVals)
	if err != nil {
		return 0, err
	}

	minDiff := 0
	for i := min; i <= max; i++ {
		diff := function(initialVals, i)
		if minDiff == 0 {
			minDiff = diff
		} else if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff, nil
}

func sumOfDiffs(initialVals []int, finalVal int) int {
	sum := 0
	for _, v := range initialVals {
		sum += utils.AbsInt(finalVal - v)
	}
	return sum
}

func progressiveSumOfDiffs(initialVals []int, finalVal int) int {
	sum := 0
	for _, v := range initialVals {
		sum += sumOfSteps(v, finalVal)
	}
	return sum
}

func sumOfSteps(val1, val2 int) int {
	diff := utils.AbsInt(val1 - val2)
	sum := 0
	for i := 1; i <= diff; i++ {
		sum += i
	}
	return sum
}
