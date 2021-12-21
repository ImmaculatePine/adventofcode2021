package day21

import "fmt"

type Player struct {
	Score    int
	Position *ListNode
}

func (p *Player) Move(steps int) {
	p.Position = p.Position.Move(steps)
}

func (p *Player) DidWin() bool {
	return p.Score >= 1000
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func (l *ListNode) Move(steps int) *ListNode {
	current := l
	for i := 0; i < steps; i++ {
		current = current.Next
	}
	return current
}

func Task1() error {
	current, other := generateBoard(8, 3)

	var rolls []int
	var totalRolls int
	var i int
	for {
		totalRolls++
		i++
		if i > 100 {
			i = 1
		}

		rolls = append(rolls, i)

		if len(rolls) == 3 {
			for _, roll := range rolls {
				current.Move(roll)
			}
			current.Score += current.Position.Value
			rolls = []int{}
			if current.DidWin() {
				fmt.Println(other.Score * totalRolls)
				return nil
			}

			*current, *other = *other, *current
		}
	}
}

func Task2() error {
	p1, p2 := generateBoard(8, 3)
	wins1, wins2 := quantumPlay(p1.Position, 0, p2.Position, 0)

	if wins1 > wins2 {
		fmt.Println(wins1)
	} else {
		fmt.Println(wins2)
	}

	return nil
}

var cache map[string][]int = make(map[string][]int)

func quantumPlay(currentPos *ListNode, currentScore int, otherPos *ListNode, otherScore int) (int, int) {
	key := fmt.Sprintf("%d-%d-%d-%d", currentPos.Value, currentScore, otherPos.Value, otherScore)
	hit, ok := cache[key]
	if ok {
		return hit[0], hit[1]
	}

	if currentScore >= 21 {
		cache[key] = []int{1, 0}
		return 1, 0
	}

	if otherScore >= 21 {
		cache[key] = []int{0, 1}
		return 0, 1
	}

	currentWins := 0
	otherWins := 0
	for _, roll := range diracDice() {
		newCurrentPos := currentPos.Move(roll)
		newCurrentScore := currentScore + newCurrentPos.Value

		ow, mw := quantumPlay(otherPos, otherScore, newCurrentPos, newCurrentScore)

		currentWins += mw
		otherWins += ow
	}

	cache[key] = []int{currentWins, otherWins}
	return currentWins, otherWins
}

func diracDice() []int {
	var rolls []int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls = append(rolls, i+j+k)
			}
		}
	}
	return rolls
}

func generateBoard(pos1, pos2 int) (*Player, *Player) {
	var head *ListNode = &ListNode{Value: -1}
	current := head
	player1 := &Player{}
	player2 := &Player{}
	for i := 1; i <= 10; i++ {
		current.Next = &ListNode{Value: i}
		current = current.Next

		if i == pos1 {
			player1.Position = current
		}

		if i == pos2 {
			player2.Position = current
		}
	}

	current.Next = head.Next

	return player1, player2
}
