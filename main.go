package main

import (
	"fmt"
	"os"

	"github.com/ImmaculatePine/adventofcode2021/day1"
	"github.com/ImmaculatePine/adventofcode2021/day2"
)

func main() {
	task := os.Args[len(os.Args)-1]
	var res string
	var err error

	switch task {
	case "1":
		res, err = day1.Task1()
	case "1+":
		res, err = day1.Task2()
	case "2":
		res, err = day2.Task1()
	case "2+":
		res, err = day2.Task2()
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
