package day5

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Point struct {
	x int
	y int
}

type Segment struct {
	p1 Point
	p2 Point
}

func (s *Segment) length() int {
	byX := math.Abs(float64(s.p1.x) - float64(s.p2.x))
	byY := math.Abs(float64(s.p1.y) - float64(s.p2.y))
	return int(math.Max(byX, byY))
}

func (s *Segment) xMultiplier() int {
	if s.p1.x < s.p2.x {
		return 1
	} else if s.p1.x > s.p2.x {
		return -1
	}
	return 0
}

func (s *Segment) yMultiplier() int {
	if s.p1.y < s.p2.y {
		return 1
	} else if s.p1.y > s.p2.y {
		return -1
	}
	return 0
}

func Task1() error {
	segments, err := readInput(func(segment Segment) bool {
		return segment.p1.x == segment.p2.x || segment.p1.y == segment.p2.y
	})
	if err != nil {
		return err
	}

	vents := drawMap(segments)
	count := sum(vents)
	fmt.Println(count)

	return nil
}

func Task2() error {
	segments, err := readInput(func(_ Segment) bool {
		return true
	})
	if err != nil {
		return nil
	}

	vents := drawMap(segments)
	count := sum(vents)
	fmt.Println(count)

	return nil
}

func drawMap(segments []Segment) map[Point]int {
	vents := make(map[Point]int)

	for _, segment := range segments {
		for i := 0; i <= segment.length(); i++ {
			point := Point{segment.p1.x + i*segment.xMultiplier(), segment.p1.y + i*segment.yMultiplier()}
			_, ok := vents[point]
			if ok {
				vents[point]++
			} else {
				vents[point] = 1
			}
		}
	}

	return vents
}

func sum(vents map[Point]int) int {
	count := 0
	for _, value := range vents {
		if value > 1 {
			count++
		}
	}
	return count
}

func readInput(predicate func(Segment) bool) ([]Segment, error) {
	lines, err := utils.ReadInputStrings("./day5/input.txt")
	if err != nil {
		return nil, err
	}

	r, err := regexp.Compile("^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$")
	if err != nil {
		return nil, err
	}

	var segments []Segment
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		if len(matches) != 5 {
			return nil, fmt.Errorf("failed to parse line %s", line)
		}

		vals, err := toInts(matches[1:])
		if err != nil {
			return nil, err
		}

		segment := Segment{
			Point{vals[0], vals[1]},
			Point{vals[2], vals[3]},
		}

		if predicate(segment) {
			segments = append(segments, segment)
		}
	}

	return segments, nil
}

func toInts(strs []string) ([]int, error) {
	var res []int
	for _, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		res = append(res, val)
	}
	return res, nil
}
