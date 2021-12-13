package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

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

func countFlashesAfterNSteps(m *utils.Matrix, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += tick(m)
	}
	return total
}

func findFirstStepWhenAllFlash(m *utils.Matrix) int {
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

func tick(m *utils.Matrix) int {
	var flashes []utils.Point
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			newVal := m.Inc(x, y)
			if newVal > 9 {
				flashes = append(flashes, utils.Point{X: x, Y: y})
			}
		}
	}

	flashes = handleFlashes(m, flashes, nil)
	for _, point := range flashes {
		m.Put(point.X, point.Y, 0)
	}

	return len(flashes)
}

func handleFlashes(m *utils.Matrix, flashes []utils.Point, acc []utils.Point) []utils.Point {
	if len(flashes) == 0 {
		return acc
	}

	acc = append(acc, flashes...)

	var newFlashes []utils.Point
	for _, point := range flashes {
		siblings := m.Siblings(point.X, point.Y)
		for sibling, val := range siblings {
			// Do not flash twice
			if val > 9 {
				continue
			}
			newVal := m.Inc(sibling.X, sibling.Y)
			if newVal > 9 {
				newFlashes = append(newFlashes, sibling)
			}
		}
	}
	return handleFlashes(m, newFlashes, acc)
}

func readInput() (*utils.Matrix, error) {
	rows, err := utils.ReadInputStrings("./day11/input.txt")
	if err != nil {
		return nil, err
	}

	matrix := utils.NewMatrix()
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
