package day20

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Matrix struct {
	data         map[Point]int
	StartX       int
	StartY       int
	Width        int
	Height       int
	DefaultValue int
}

type Point struct {
	X int
	Y int
}

func NewMatrix() *Matrix {
	return &Matrix{
		data:         make(map[Point]int),
		StartX:       0,
		StartY:       0,
		Width:        0,
		Height:       0,
		DefaultValue: 0,
	}
}

func (m *Matrix) Put(x, y, val int) {
	point := Point{x, y}
	m.data[point] = val
	if x < m.StartX {
		m.StartX = x
	}
	if y < m.StartY {
		m.StartY = y
	}
	if x+1 > m.Width {
		m.Width = x + 1
	}
	if y+1 > m.Height {
		m.Height = y + 1
	}
}

func (m *Matrix) Get(x, y int) int {
	point := Point{x, y}
	val, ok := m.data[point]
	if !ok {
		return m.DefaultValue
	}
	return val
}

func (m *Matrix) RemoveBorders() {
	for x := m.StartX; x < m.Width; x++ {
		delete(m.data, Point{x, m.StartY})
		delete(m.data, Point{x, m.Height - 1})
	}

	for y := m.StartY; y < m.Height; y++ {
		delete(m.data, Point{m.StartX, y})
		delete(m.data, Point{m.Width - 1, y})
	}

	m.StartX++
	m.StartY++
	m.Width--
	m.Height--
}

func (m *Matrix) Square(x, y int) []int {
	points := []Point{
		// Top row
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		// Middle row
		{x - 1, y},
		{x, y},
		{x + 1, y},
		// Bottom row
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}

	var result []int
	for _, point := range points {
		result = append(result, m.Get(point.X, point.Y))
	}
	return result
}

func (m *Matrix) Count() int {
	var sum int
	for _, val := range m.data {
		if val > 0 {
			sum++
		}
	}
	return sum
}

func (m *Matrix) Print() error {
	for y := m.StartY; y < m.Height; y++ {
		for x := m.StartX; x < m.Width; x++ {
			val := m.Get(x, y)
			if val > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	return nil
}

func Task1() error {
	algorithm, input, err := readInput()
	if err != nil {
		return err
	}

	output, err := enhanceTimes(input, algorithm, 2)
	if err != nil {
		return err
	}

	fmt.Println(output.Count())

	return nil
}

func Task2() error {
	algorithm, input, err := readInput()
	if err != nil {
		return err
	}

	output, err := enhanceTimes(input, algorithm, 50)
	if err != nil {
		return err
	}

	fmt.Println(output.Count())

	return nil
}

func enhanceTimes(input *Matrix, algorithm []int, times int) (*Matrix, error) {
	var output *Matrix = input
	var err error
	for i := 0; i < times; i++ {
		output, err = enhance(output, algorithm)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func enhance(input *Matrix, algorithm []int) (*Matrix, error) {
	output := NewMatrix()
	for x := input.StartX - 3; x < input.Width+3; x++ {
		for y := input.StartY - 3; y < input.Height+3; y++ {
			data := input.Square(x, y)
			val, err := decode(data, algorithm)
			if err != nil {
				return nil, err
			}
			output.Put(x, y, val)
		}
	}
	removeBorder(output)
	return output, nil
}

func decode(data []int, algorithm []int) (int, error) {
	var binary string
	for _, bit := range data {
		if bit > 0 {
			binary += "1"
		} else {
			binary += "0"
		}
	}

	index, err := strconv.ParseUint(binary, 2, 64)
	if err != nil {
		return 0, err
	}

	if int(index) > len(algorithm)-1 {
		return -1, fmt.Errorf("unexpected index %d", index)
	}

	val := algorithm[index]
	return val, nil
}

func removeBorder(matrix *Matrix) {
	topLeft := matrix.Get(matrix.StartX, matrix.StartY)
	for x := matrix.StartX; x < matrix.Width; x++ {
		if matrix.Get(x, matrix.StartY) != topLeft {
			return
		}
		if matrix.Get(x, matrix.Height-1) != topLeft {
			return
		}
	}

	for y := matrix.StartY; y < matrix.Height; y++ {
		if matrix.Get(matrix.StartX, y) != topLeft {
			return
		}
		if matrix.Get(matrix.Height-1, y) != topLeft {
			return
		}
	}

	// Didn't return yet? There is a border, let's remove it and repeat.
	matrix.DefaultValue = topLeft
	matrix.RemoveBorders()
	removeBorder(matrix)
}

func readInput() ([]int, *Matrix, error) {
	file, err := os.Open("./day20/input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read algorithm from the first line
	scanner.Scan()
	var algorithm []int
	for _, r := range scanner.Text() {
		val, err := charToInt(r)
		if err != nil {
			return nil, nil, err
		}
		algorithm = append(algorithm, val)
	}

	// Skip the empty line
	scanner.Scan()

	// Read input image
	matrix := NewMatrix()
	y := 0
	for scanner.Scan() {
		for x, r := range scanner.Text() {
			val, err := charToInt(r)
			if err != nil {
				return nil, nil, err
			}
			matrix.Put(x, y, val)
		}
		y++
	}

	return algorithm, matrix, scanner.Err()
}

func charToInt(char rune) (int, error) {
	switch char {
	case '.':
		return 0, nil
	case '#':
		return 1, nil
	default:
		return -1, fmt.Errorf("unexpected character %c", char)
	}
}
