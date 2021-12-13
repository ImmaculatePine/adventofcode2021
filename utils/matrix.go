package utils

import "fmt"

type Matrix struct {
	data   map[Point]int
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
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
		val, err := m.Get(point.X, point.Y)
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
