package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ImmaculatePine/adventofcode2021/day1"
	"github.com/ImmaculatePine/adventofcode2021/day2"
	"github.com/ImmaculatePine/adventofcode2021/day3"
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
	default:
		err = fmt.Errorf("unknown task %s", task)
	}

	if err != nil {
		log.Fatalf("%v", err)
	}
}
