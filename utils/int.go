package utils

import "fmt"

func FindMinInt(vals []int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("can't find min value of empty slice")
	}

	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func FindMaxInt(vals []int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("can't find max value of empty slice")
	}

	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func AbsInt(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}
