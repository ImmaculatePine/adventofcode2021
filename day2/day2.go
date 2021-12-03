package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Command struct {
	direction string
	value     int
}

func readCommands() ([]*Command, error) {
	input, err := utils.ReadInputStrings("./day2/input.txt")
	if err != nil {
		return nil, err
	}

	var commands []*Command
	for _, line := range input {
		cmd, err := parseCommand(line)
		if err != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}

	return commands, nil
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

	return &Command{fields[0], val}, nil
}

func Task1() error {
	commands, err := readCommands()
	if err != nil {
		return err
	}

	result, err := task1(commands)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func task1(commands []*Command) (int, error) {
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
		default:
			return 0, fmt.Errorf("unknown command %s", cmd.direction)
		}
	}

	return pos * depth, nil
}

func Task2() error {
	commands, err := readCommands()
	if err != nil {
		return err
	}

	result, err := task2(commands)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func task2(commands []*Command) (int, error) {
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
		default:
			return 0, fmt.Errorf("unknown command %s", cmd.direction)
		}
	}

	return pos * depth, nil
}
