package day12

import (
	"fmt"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Graph struct {
	Edges map[string][]string
	paths []Path
}

type Path []string

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]string),
		paths: []Path{},
	}
}

func (g *Graph) AddEdge(from, to string) {
	edges, ok := g.Edges[from]
	if ok {
		g.Edges[from] = append(edges, to)
	} else {
		g.Edges[from] = []string{to}
	}
}

func (g *Graph) FindAllPaths(start, goal string, isAllowed func(string, Path) bool) []Path {
	g.findAllPaths(start, goal, Path{start}, isAllowed)
	return g.paths
}

func (g *Graph) findAllPaths(node, goal string, currentPath Path, isAllowed func(string, Path) bool) {
	if node == goal {
		g.paths = append(g.paths, currentPath)
		return
	}

	transitions, ok := g.Edges[node]
	if !ok {
		return
	}

	for _, transition := range transitions {
		if isAllowed(transition, currentPath) {
			g.findAllPaths(transition, goal, append(currentPath, transition), isAllowed)
		}
	}
}

func Task1() error {
	graph, err := readInput()
	if err != nil {
		return err
	}

	paths := graph.FindAllPaths("start", "end", func(node string, path Path) bool {
		if !isSmall(node) {
			return true
		}
		for _, n := range path {
			if n == node {
				return false
			}
		}
		return true
	})

	fmt.Println(len(paths))

	return nil
}

func Task2() error {
	graph, err := readInput()
	if err != nil {
		return err
	}

	paths := graph.FindAllPaths("start", "end", func(node string, path Path) bool {
		if !isSmall(node) {
			return true
		}

		counters := make(map[string]int)
		didVisitTwice := false
		for _, n := range path {
			if isSmall(n) {
				counters[n] = counters[n] + 1
				if counters[n] > 1 {
					didVisitTwice = true
				}
			}
		}

		switch counters[node] {
		case 0:
			return true
		case 1:
			return !didVisitTwice && node != "start" && node != "end"
		default:
			return false
		}
	})

	fmt.Println(len(paths))

	return nil
}

func isSmall(node string) bool {
	return node == strings.ToLower(node)
}

func readInput() (*Graph, error) {
	lines, err := utils.ReadInputStrings("./day12/input.txt")
	if err != nil {
		return nil, err
	}

	graph := NewGraph()
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected 2 parts, got %d", len(parts))
		}
		graph.AddEdge(parts[0], parts[1])
		graph.AddEdge(parts[1], parts[0])
	}
	return graph, nil
}
