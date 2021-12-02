package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Command struct {
	direction string
	value     int
}

func main() {
	input, err := utils.ReadInputStrings("./day2/input.txt")
	if err != nil {
		log.Fatalf("failed to read input, %v", err)
	}

	var commands []*Command
	for _, line := range input {
		cmd, err := parseCommand(line)
		if err != nil {
			log.Fatalf("failed to parse input, %v", err)
		}
		commands = append(commands, cmd)
	}

	switch os.Args[len(os.Args)-1] {
	case "1":
		result, err := Task1(commands)
		if err != nil {
			log.Fatalf("task 1 failed, %v", err)
		}
		fmt.Println(result)
	case "2":
		result, err := Task2(commands)
		if err != nil {
			log.Fatalf("task 2 failed, %v", err)
		}
		fmt.Println(result)
	default:
		log.Fatalf("unknown task %s", os.Args[len(os.Args)-1])
	}
}

func parseCommand(str string) (*Command, error) {
	fields := strings.Fields(str)
	if len(fields) != 2 {
		return nil, fmt.Errorf("unknown command %s", str)
	}

	if len(fields) != 2 {
		return nil, fmt.Errorf("failed to parse command %s", str)
	}

	val, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse value in %s", str)
	}

	switch fields[0] {
	case "forward":
		return &Command{"forward", val}, nil
	case "up":
		return &Command{"up", val}, nil
	case "down":
		return &Command{"down", val}, nil
	default:
		return nil, fmt.Errorf("unknown command %s", fields[0])
	}
}

func Task1(commands []*Command) (int, error) {
	pos := 0
	depth := 0

	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			pos += cmd.value
		case "up":
			depth -= cmd.value
		case "down":
			depth += cmd.value
		}
	}

	return pos * depth, nil
}

func Task2(commands []*Command) (int, error) {
	pos := 0
	depth := 0
	aim := 0

	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			pos += cmd.value
			depth += aim * cmd.value
		case "up":
			aim -= cmd.value
		case "down":
			aim += cmd.value
		}
	}

	return pos * depth, nil
}
