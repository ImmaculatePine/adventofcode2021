package day18

import (
	"fmt"
	"math"
	"strconv"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func (n *Node) IsRegular() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) FindClosestLeftRegular() *Node {
	prev := n
	for current := n.parent; current != nil; {
		if current.right == prev {
			prev = current
			current = current.left
			continue
		}
		if current.left == prev {
			prev = current
			current = current.parent
			continue
		}
		if current.IsRegular() {
			return current
		}
		current = current.right
	}

	return nil
}

func (n *Node) FindClosestRightRegular() *Node {
	prev := n
	for current := n.parent; current != nil; {
		if current.left == prev {
			prev = current
			current = current.right
			continue
		}
		if current.right == prev {
			prev = current
			current = current.parent
			continue
		}
		if current.IsRegular() {
			return current
		}
		current = current.left
	}

	return nil
}

func (n *Node) Magnitude() int {
	if n.IsRegular() {
		return n.value
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n *Node) ToString() string {
	if n == nil {
		return ""
	}

	if n.left != nil && n.right != nil {
		return fmt.Sprintf("[%s,%s]", n.left.ToString(), n.right.ToString())
	}

	return fmt.Sprintf("%d", n.value)
}

func Task1() error {
	lines, err := utils.ReadInputStrings("./day18/input.txt")
	if err != nil {
		return err
	}

	var root *Node
	for _, line := range lines {
		tree, err := parseTree(line)
		if err != nil {
			return err
		}

		if root == nil {
			root = tree
			continue
		}

		root = add(root, tree)
		reduce(root)
	}

	fmt.Println(root.Magnitude())

	return nil
}

func Task2() error {
	lines, err := utils.ReadInputStrings("./day18/input.txt")
	if err != nil {
		return err
	}

	maxMagnitude := 0
	for _, a := range lines {
		for _, b := range lines {
			if a == b {
				continue
			}

			treeA, err := parseTree(a)
			if err != nil {
				return err
			}

			treeB, err := parseTree(b)
			if err != nil {
				return err
			}

			sum := add(treeA, treeB)
			reduce(sum)
			magnitude := sum.Magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}

	fmt.Println(maxMagnitude)

	return nil
}

func add(left, right *Node) *Node {
	root := &Node{left: left, right: right}
	left.parent = root
	right.parent = root
	return root
}

func explode(node *Node) bool {
	explodedNode := findNodeToExplode(node)
	if explodedNode == nil {
		return false
	}

	leftDiff := explodedNode.left.value
	rightDiff := explodedNode.right.value

	explodedNode.left = nil
	explodedNode.right = nil

	closestLeft := explodedNode.FindClosestLeftRegular()
	if closestLeft != nil {
		closestLeft.value += leftDiff
	}

	closestRight := explodedNode.FindClosestRightRegular()
	if closestRight != nil {
		closestRight.value += rightDiff
	}

	return true
}

func split(node *Node) bool {
	splitNode := findNodeToSplit(node)
	if splitNode == nil {
		return false
	}

	leftValue := math.Floor(float64(splitNode.value) / 2)
	rightValue := math.Ceil(float64(splitNode.value) / 2)

	splitNode.value = 0
	splitNode.left = &Node{
		value:  int(leftValue),
		parent: splitNode,
	}
	splitNode.right = &Node{
		value:  int(rightValue),
		parent: splitNode,
	}

	return true
}

func reduce(node *Node) {
	if explode(node) {
		reduce(node)
		return
	}
	if split(node) {
		reduce(node)
		return
	}
}

func findNodeToSplit(tree *Node) *Node {
	if tree == nil {
		return nil
	}

	if tree.IsRegular() && tree.value >= 10 {
		return tree
	}

	node := findNodeToSplit(tree.left)
	if node != nil {
		return node
	}

	node = findNodeToSplit(tree.right)
	if node != nil {
		return node
	}

	return nil
}

func findNodeToExplode(tree *Node) *Node {
	return doFindNodesToExplode(tree, 0)
}

func doFindNodesToExplode(tree *Node, level int) *Node {
	if tree.IsRegular() {
		return nil
	}

	if level >= 4 {
		return tree
	}

	node := doFindNodesToExplode(tree.left, level+1)
	if node != nil {
		return node
	}

	node = doFindNodesToExplode(tree.right, level+1)
	if node != nil {
		return node
	}

	return nil
}

func parseTree(input string) (*Node, error) {
	node := &Node{}
	root := node
	buffer := ""
	for _, r := range input {
		switch r {
		case '[':
			node.left = &Node{parent: node}
			node.right = &Node{parent: node}
			node = node.left
		case ',':
			if len(buffer) > 0 {
				value, err := strconv.Atoi(buffer)
				if err != nil {
					return nil, fmt.Errorf("unexpected number %s in %s", buffer, input)
				}
				node.value = value
				buffer = ""
			}
			node = node.parent.right
		case ']':
			if len(buffer) > 0 {
				value, err := strconv.Atoi(buffer)
				if err != nil {
					return nil, fmt.Errorf("unexpected number %s in %s", buffer, input)
				}
				node.value = value
				buffer = ""
			}
			node = node.parent
		default:
			buffer += string(r)
		}
	}
	return root, nil
}
