package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ImmaculatePine/adventofcode2021/day1"
	"github.com/ImmaculatePine/adventofcode2021/day2"
	"github.com/ImmaculatePine/adventofcode2021/day3"
	"github.com/ImmaculatePine/adventofcode2021/day4"
	"github.com/ImmaculatePine/adventofcode2021/day5"
	"github.com/ImmaculatePine/adventofcode2021/day6"
	"github.com/ImmaculatePine/adventofcode2021/day7"
	"github.com/ImmaculatePine/adventofcode2021/day8"
	"github.com/ImmaculatePine/adventofcode2021/day9"
)

func main() {
	task := os.Args[len(os.Args)-1]
	var err error

	switch task {
	case "1":
		err = day1.Task1()
	case "1+":
		err = day1.Task2()
	case "2":
		err = day2.Task1()
	case "2+":
		err = day2.Task2()
	case "3":
		err = day3.Task1()
	case "3+":
		err = day3.Task2()
	case "4":
		err = day4.Task1()
	case "4+":
		err = day4.Task2()
	case "5":
		err = day5.Task1()
	case "5+":
		err = day5.Task2()
	case "6":
		err = day6.Task1()
	case "6+":
		err = day6.Task2()
	case "7":
		err = day7.Task1()
	case "7+":
		err = day7.Task2()
	case "8":
		err = day8.Task1()
	case "8+":
		err = day8.Task2()
	case "9":
		err = day9.Task1()
	case "9+":
		err = day9.Task2()
	default:
		err = fmt.Errorf("unknown task %s", task)
	}

	if err != nil {
		log.Fatalf("%v", err)
	}
}
