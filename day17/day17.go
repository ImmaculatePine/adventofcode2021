package day17

import (
	"fmt"
)

func Task1() error {
	result, _ := solve()
	fmt.Println(result)
	return nil
}

func Task2() error {
	_, result := solve()
	fmt.Println(result)
	return nil
}

func solve() (int, int) {
	leftX, rightX, bottomY, topY := 155, 182, -117, -67

	finalMaxY := bottomY
	matchesCount := 0
	for initialX := -1000; initialX < 1000; initialX++ {
		for initialY := bottomY; initialY < -bottomY; initialY++ {
			x := 0
			y := 0
			maxY := 0
			vx := initialX
			vy := initialY
			for {
				x += vx
				y += vy
				if vx > 0 {
					vx--
				} else if vx < 0 {
					vx++
				}
				vy--

				if y > maxY {
					maxY = y
				}

				if y < bottomY {
					break
				}

				if x >= leftX && x <= rightX && y >= bottomY && y <= topY {
					matchesCount++
					if maxY > finalMaxY {
						finalMaxY = maxY
					}
					break
				}
			}
		}
	}

	return finalMaxY, matchesCount
}
