package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Matrix struct {
	data   map[Point]int
	Width  int
	Height int
}

type Point struct {
	x int
	y int
}

func NewMatrix() *Matrix {
	return &Matrix{
		data:   make(map[Point]int),
		Width:  0,
		Height: 0,
	}
}

func (m *Matrix) Put(x, y, val int) {
	point := Point{x, y}
	m.data[point] = val
	if x+1 > m.Width {
		m.Width = x + 1
	}
	if y+1 > m.Height {
		m.Height = y + 1
	}
}

func (m *Matrix) Get(x, y int) (int, error) {
	point := Point{x, y}
	val, ok := m.data[point]
	if !ok {
		return 0, fmt.Errorf("(%d, %d) not found", x, y)
	}
	return val, nil
}

func (m *Matrix) Inc(x, y int) int {
	point := Point{x, y}
	val, ok := m.data[point]
	if ok {
		m.data[point]++
		return val + 1
	}
	m.data[point] = 1
	return 1
}

func (m *Matrix) Siblings(x, y int) map[Point]int {
	points := []Point{
		// Top row
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		// Middle row
		{x - 1, y},
		{x + 1, y},
		// Bottom row
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}

	result := make(map[Point]int)
	for _, point := range points {
		val, err := m.Get(point.x, point.y)
		if err == nil {
			result[point] = val
		}
	}
	return result
}

func (m *Matrix) Print() error {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			val, err := m.Get(x, y)
			if err != nil {
				return err
			}
			fmt.Printf("%d", val)
		}
		fmt.Printf("\n")
	}
	return nil
}

func Task1() error {
	matrix, err := readInput()
	if err != nil {
		return err
	}

	fmt.Println(countFlashesAfterNSteps(matrix, 100))

	return nil
}

func Task2() error {
	matrix, err := readInput()
	if err != nil {
		return err
	}

	fmt.Println(findFirstStepWhenAllFlash(matrix))

	return nil
}

func countFlashesAfterNSteps(m *Matrix, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += tick(m)
	}
	return total
}

func findFirstStepWhenAllFlash(m *Matrix) int {
	step := 0
	size := m.Width * m.Height
	for {
		step++
		flashes := tick(m)
		if flashes == size {
			return step
		}
	}
}

func tick(m *Matrix) int {
	var flashes []Point
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			newVal := m.Inc(x, y)
			if newVal > 9 {
				flashes = append(flashes, Point{x, y})
			}
		}
	}

	flashes = handleFlashes(m, flashes, nil)
	for _, point := range flashes {
		m.Put(point.x, point.y, 0)
	}

	return len(flashes)
}

func handleFlashes(m *Matrix, flashes []Point, acc []Point) []Point {
	if len(flashes) == 0 {
		return acc
	}

	acc = append(acc, flashes...)

	var newFlashes []Point
	for _, point := range flashes {
		siblings := m.Siblings(point.x, point.y)
		for sibling, val := range siblings {
			// Do not flash twice
			if val > 9 {
				continue
			}
			newVal := m.Inc(sibling.x, sibling.y)
			if newVal > 9 {
				newFlashes = append(newFlashes, sibling)
			}
		}
	}
	return handleFlashes(m, newFlashes, acc)
}

func readInput() (*Matrix, error) {
	rows, err := utils.ReadInputStrings("./day11/input.txt")
	if err != nil {
		return nil, err
	}

	matrix := NewMatrix()
	for y, row := range rows {
		cols := strings.Split(row, "")
		for x, col := range cols {
			val, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			matrix.Put(x, y, val)
		}
	}
	return matrix, nil
}
