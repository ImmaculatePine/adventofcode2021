package day9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Matrix map[int]map[int]int

type Point struct {
	x   int
	y   int
	val int
}

func (m Matrix) put(x, y, val int) {
	_, ok := m[y]
	if !ok {
		m[y] = make(map[int]int)
	}
	m[y][x] = val
}

func (m Matrix) get(x, y int) (int, error) {
	row, ok := m[y]
	if !ok {
		return 0, fmt.Errorf("(%d, %d) not found", x, y)
	}
	val, ok := row[x]
	if !ok {
		return 0, fmt.Errorf("(%d, %d) not found", x, y)
	}
	return val, nil
}

func Task1() error {
	matrix, err := readInput()
	if err != nil {
		return nil
	}

	lows, err := findLows(matrix)
	if err != nil {
		return nil
	}

	sum := 0
	for _, point := range lows {
		sum += point.val + 1
	}

	fmt.Println(sum)

	return nil
}

func Task2() error {
	matrix, err := readInput()
	if err != nil {
		return nil
	}

	lows, err := findLows(matrix)
	if err != nil {
		return nil
	}

	// Key is the lowest point, value is the bassin size
	basins := make(map[Point]int)
	for _, point := range lows {
		size := basinSize(matrix, point)
		basins[point] = size
	}

	sizes := make([]int, 0, len(basins))
	for _, size := range basins {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	fmt.Println(sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3])

	return nil
}

func basinSize(matrix Matrix, low Point) int {
	return traverseBasin(matrix, low, 0)
}

func traverseBasin(matrix Matrix, point Point, acc int) int {
	siblings := []Point{}
	up, err := matrix.get(point.x, point.y-1)
	if err == nil && up > point.val && up != 9 {
		siblings = append(siblings, Point{point.x, point.y - 1, up})
	}

	down, err := matrix.get(point.x, point.y+1)
	if err == nil && down > point.val && down != 9 {
		siblings = append(siblings, Point{point.x, point.y + 1, down})
	}

	left, err := matrix.get(point.x-1, point.y)
	if err == nil && left > point.val && left != 9 {
		siblings = append(siblings, Point{point.x - 1, point.y, left})
	}

	right, err := matrix.get(point.x+1, point.y)
	if err == nil && right > point.val && right != 9 {
		siblings = append(siblings, Point{point.x + 1, point.y, right})
	}

	matrix.put(point.x, point.y, 9)

	for _, sibling := range siblings {
		matrix.put(sibling.x, sibling.y, 9)
	}

	for _, sibling := range siblings {
		acc += traverseBasin(matrix, sibling, 0)
	}

	return acc + 1
}

func findLows(matrix Matrix) ([]Point, error) {
	var points []Point
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			val, err := matrix.get(x, y)
			if err != nil {
				return nil, err
			}

			up, err := matrix.get(x, y-1)
			if err == nil && val >= up {
				continue
			}

			down, err := matrix.get(x, y+1)
			if err == nil && val >= down {
				continue
			}

			left, err := matrix.get(x-1, y)
			if err == nil && val >= left {
				continue
			}

			right, err := matrix.get(x+1, y)
			if err == nil && val >= right {
				continue
			}

			points = append(points, Point{x, y, val})
		}
	}
	return points, nil
}

func readInput() (Matrix, error) {
	lines, err := utils.ReadInputStrings("./day9/input.txt")
	if err != nil {
		return nil, err
	}

	matrix := make(Matrix)
	for y, line := range lines {
		parts := strings.Split(line, "")
		for x, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			matrix.put(x, y, val)
		}
	}
	return matrix, nil
}
