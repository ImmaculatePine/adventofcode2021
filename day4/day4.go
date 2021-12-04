package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Numbers = []int
type UnmarkedCount = int

type Board struct {
	items  map[int]*BoardItem
	rows   map[int]UnmarkedCount
	cols   map[int]UnmarkedCount
	didWin bool
}

type BoardItem struct {
	row      int
	col      int
	isMarked bool
}

func Task1() error {
	numbers, boards, err := readInput()
	if err != nil {
		return err
	}

	for _, number := range numbers {
		for _, board := range boards {
			didWin := playBoard(board, number)
			if didWin {
				result := sumUnmarked(board) * number
				fmt.Println(result)
				return nil
			}
		}
	}

	return fmt.Errorf("no winners")
}

func Task2() error {
	numbers, boards, err := readInput()
	if err != nil {
		return err
	}

	var lastWonBoard *Board
	var lastWonNumber int

	for _, number := range numbers {
		for _, board := range boards {
			didWin := playBoard(board, number)
			if didWin {
				board.didWin = true
				lastWonBoard = board
				lastWonNumber = number
			}
		}
	}

	if lastWonBoard == nil {
		return fmt.Errorf("no winners")
	}

	result := sumUnmarked(lastWonBoard) * lastWonNumber
	fmt.Println(result)

	return nil
}

func playBoard(board *Board, number int) bool {
	// Skip already board that already won
	if board.didWin {
		return false
	}

	item := board.items[number]
	if item == nil {
		return false
	}

	if item.isMarked {
		return false
	}

	item.isMarked = true
	board.rows[item.row]--
	board.cols[item.col]--

	return board.rows[item.row] == 0 || board.cols[item.col] == 0
}

func sumUnmarked(board *Board) (sum int) {
	for number, item := range board.items {
		if !item.isMarked {
			sum += number
		}
	}
	return
}

func readInput() (Numbers, []*Board, error) {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read numbers from the first line
	scanner.Scan()
	numberStrings := strings.Split(scanner.Text(), ",")
	var numbers Numbers
	for _, str := range numberStrings {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, nil, err
		}
		numbers = append(numbers, val)
	}

	// Start reading boards
	var boards []*Board
	var board *Board
	var currentRow int

	for scanner.Scan() {
		// An empty line is a marker of a new board
		if scanner.Text() == "" {
			// Finalize the previous board first
			if board != nil {
				colSize := len(board.rows)
				for colIdx := 0; colIdx < currentRow; colIdx++ {
					board.cols[colIdx] = colSize
				}
				boards = append(boards, board)
			}

			// And initialize a new one
			board = &Board{
				items: make(map[int]*BoardItem),
				rows:  make(map[int]UnmarkedCount),
				cols:  make(map[int]UnmarkedCount),
			}
			currentRow = 0
			continue
		}

		row := strings.Fields(scanner.Text())
		board.rows[currentRow] = len(row)
		for currentCol, number := range row {
			val, err := strconv.Atoi(number)
			if err != nil {
				return nil, nil, err
			}

			board.items[val] = &BoardItem{
				row:      currentRow,
				col:      currentCol,
				isMarked: false,
			}
		}

		currentRow++
	}

	// Finalize and push the last parsed board to the list
	colSize := len(board.rows)
	for colIdx := 0; colIdx < currentRow; colIdx++ {
		board.cols[colIdx] = colSize
	}
	boards = append(boards, board)

	return numbers, boards, scanner.Err()
}
