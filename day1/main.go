package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

func main() {
	input, err := utils.ReadInputInts("./day1/input.txt")
	if err != nil {
		log.Fatalf("failed to read input, %v", err)
	}

	switch os.Args[len(os.Args)-1] {
	case "1":
		result, err := GetNumberOfIncreases(input)
		if err != nil {
			log.Fatalf("task 1 failed, %v", err)
		}
		fmt.Println(result)
	case "2":
		result, err := GetNumberOfWindowIncreases(input)
		if err != nil {
			log.Fatalf("task 2 failed, %v", err)
		}
		fmt.Println(result)
	default:
		log.Fatalf("unknown task %s", os.Args[len(os.Args)-1])
	}
}

func GetNumberOfIncreases(input []int) (int, error) {
	result := 0
	last := input[0]

	for _, v := range input {
		if v > last {
			result++
		}
		last = v
	}

	return result, nil
}

func GetNumberOfWindowIncreases(input []int) (int, error) {
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

	return result, nil
}

func sum(arr []int) (res int) {
	for _, v := range arr {
		res += v
	}
	return
}
