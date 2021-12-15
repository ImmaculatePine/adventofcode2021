package day15

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Node struct {
	Weight   int
	Edges    []*Node
	priority int
	index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Node, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Dijkstra(start, end *Node) (int, error) {
	distances := make(map[*Node]int)
	queue := PriorityQueue{}
	heap.Init(&queue)

	distances[start] = 0
	start.priority = 0
	heap.Push(&queue, start)

	for _, node := range start.Edges {
		distances[node] = math.MaxInt
		node.priority = math.MaxInt
		heap.Push(&queue, node)
	}

	for queue.Len() > 0 {
		node := heap.Pop(&queue).(*Node)

		for _, edge := range node.Edges {
			newPathLength := getDistance(distances, node) + edge.Weight
			oldPathLength := getDistance(distances, edge)
			if newPathLength < oldPathLength {
				distances[edge] = newPathLength
				if edge.priority > 0 {
					queue.update(edge, newPathLength)
				} else {
					edge.priority = newPathLength
					heap.Push(&queue, edge)
				}
			}
		}
	}

	res, ok := distances[end]
	if !ok {
		return 0, fmt.Errorf("start and end nodes are not connected")
	}

	return res, nil
}

func getDistance(distances map[*Node]int, node *Node) int {
	val, ok := distances[node]
	if ok {
		return val
	}
	return math.MaxInt
}

func Task1() error {
	start, end, err := readInput(1)
	if err != nil {
		return err
	}

	res, err := Dijkstra(start, end)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func Task2() error {
	start, end, err := readInput(5)
	if err != nil {
		return err
	}

	res, err := Dijkstra(start, end)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func readInput(repeat int) (*Node, *Node, error) {
	rows, err := utils.ReadInputStrings("./day15/input.txt")
	if err != nil {
		return nil, nil, err
	}

	var nodes map[utils.Point]*Node = make(map[utils.Point]*Node)
	for i := 0; i < repeat; i++ {
		for y, row := range rows {
			cols := strings.Split(row, "")
			for j := 0; j < repeat; j++ {
				for x, col := range cols {
					val, err := strconv.Atoi(col)
					if err != nil {
						return nil, nil, err
					}
					weight := val + i + j
					if weight > 9 {
						weight -= 9
					}
					node := &Node{Weight: weight}
					nodes[utils.Point{X: j*len(rows[0]) + x, Y: i*len(rows) + y}] = node
				}
			}
		}
	}

	for x := 0; x < len(rows[0])*repeat; x++ {
		for y := 0; y < len(rows)*repeat; y++ {
			node, ok := nodes[utils.Point{X: x, Y: y}]
			if !ok {
				return nil, nil, fmt.Errorf("missing node (%d,%d)", x, y)
			}

			points := []utils.Point{
				// Top row
				{X: x, Y: y - 1},
				// Middle row
				{X: x - 1, Y: y},
				{X: x + 1, Y: y},
				// Bottom row
				{X: x, Y: y + 1},
			}

			for _, point := range points {
				edge, ok := nodes[point]
				if ok {
					node.Edges = append(node.Edges, edge)
				}
			}
		}
	}

	start := nodes[utils.Point{X: 0, Y: 0}]
	end := nodes[utils.Point{X: len(rows[0])*repeat - 1, Y: len(rows)*repeat - 1}]
	return start, end, nil
}
