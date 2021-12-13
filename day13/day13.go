package day13

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Paper struct {
	data   map[utils.Point]bool
	Width  int
	Height int
}

func NewPaper() *Paper {
	return &Paper{
		data: make(map[utils.Point]bool),
	}
}

func (p *Paper) Add(x, y int) {
	point := utils.Point{X: x, Y: y}
	p.data[point] = true
	if x+1 > p.Width {
		p.Width = x + 1
	}
	if y+1 > p.Height {
		p.Height = y + 1
	}
}

func (p *Paper) DotsCount() int {
	sum := 0
	for _, v := range p.data {
		if v {
			sum++
		}
	}
	return sum
}

func (p *Paper) Print() {
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			val := p.data[utils.Point{X: x, Y: y}]
			if val {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (p *Paper) Fold(fold Fold) error {
	switch fold.axis {
	case "x":
		for point := range p.data {
			if point.X == fold.value {
				p.data[point] = false
			} else if point.X > fold.value {
				newPoint := utils.Point{X: p.Width - point.X - 1, Y: point.Y}
				p.data[point] = false
				p.data[newPoint] = true
			}
		}
		p.Width = (p.Width - 1) / 2
		return nil
	case "y":
		for point := range p.data {
			if point.Y == fold.value {
				p.data[point] = false
			} else if point.Y > fold.value {
				newPoint := utils.Point{X: point.X, Y: p.Height - point.Y - 1}
				p.data[point] = false
				p.data[newPoint] = true
			}
		}
		p.Height = (p.Height - 1) / 2
		return nil
	default:
		return fmt.Errorf("uknown fold axis %s", fold.axis)
	}
}

type Fold struct {
	axis  string
	value int
}

func Task1() error {
	paper, folds, err := readInput()
	if err != nil {
		return err
	}

	paper.Fold(folds[0])
	fmt.Println(paper.DotsCount())

	return nil
}

func Task2() error {
	paper, folds, err := readInput()
	if err != nil {
		return err
	}

	for _, fold := range folds {
		paper.Fold(fold)
	}

	paper.Print()

	return nil
}

func readInput() (*Paper, []Fold, error) {
	file, err := os.Open("./day13/input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan dots matrix
	paper := NewPaper()
	for scanner.Scan() && scanner.Text() != "" {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("expected 2 comma-separated values, got %s", scanner.Text())
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		paper.Add(x, y)
	}

	// Fold regexp
	r, err := regexp.Compile("^fold along (x|y)=([0-9]+)$")
	if err != nil {
		return nil, nil, err
	}

	// Scan fold rules
	folds := []Fold{}
	for scanner.Scan() {
		matches := r.FindStringSubmatch(scanner.Text())
		if len(matches) != 3 {
			return nil, nil, fmt.Errorf("failed to parse line %s", scanner.Text())
		}

		val, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, nil, err
		}

		folds = append(folds, Fold{matches[1], val})
	}

	return paper, folds, nil
}
