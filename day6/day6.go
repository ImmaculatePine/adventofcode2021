package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Timers = map[int]int

func Task1() error {
	timers, err := readInput()
	if err != nil {
		return err
	}

	timers = simulate(timers, 80)
	fmt.Println(size(timers))

	return nil
}

func Task2() error {
	timers, err := readInput()
	if err != nil {
		return err
	}

	timers = simulate(timers, 256)
	fmt.Println(size(timers))
	return nil
}

func size(timers Timers) int {
	sum := 0
	for _, v := range timers {
		sum += v
	}
	return sum
}

func simulate(timers Timers, days int) Timers {
	for i := 0; i < days; i++ {
		timers = tick(timers)
	}
	return timers
}

func tick(timers Timers) Timers {
	newTimers := make(map[int]int)

	resetCount := 0

	for timer, count := range timers {
		if timer > 0 {
			newTimers[timer-1] = count
		} else if timer == 0 {
			newTimers[8] = count
			resetCount = count
		}
	}

	newTimers[6] = newTimers[6] + resetCount

	return newTimers
}

func readInput() (Timers, error) {
	input, err := utils.ReadInputStrings("./day6/input.txt")
	if err != nil {
		return nil, err
	}

	parts := strings.Split(input[0], ",")

	timers := make(map[int]int)
	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		timers[val] = timers[val] + 1
	}

	return timers, nil
}
