package solver

import (
	"fmt"
)

type solver struct {
	sudoku [][]int
}

func NewSolver(sudoku [][]int) *solver {
	return &solver{sudoku: sudoku}
}

func (s *solver) Solve() [][]int {
	filled := make([]bool, 81)
	placement := initPlacement(s.sudoku)

	for !isComplete(filled) {
		for i, row := range s.sudoku {
			for j, cell := range row {
				if cell > 0 {
					filled[i*9+j] = true
					continue
				}
				possible := getPossibleByElimination(i, j, s.sudoku)
				cellNode := placement[i*9+j]
				cellNode.possible = possible
				if len(possible) == 1 {
					s.sudoku[i][j] = possible[0]
					filled[i*9+j] = true
					continue
				}
				for _, p := range possible {
					if isSingleOccurrenceOfNumber(i, j, p, placement) {
						cellNode.possible = []int{p}
						s.sudoku[i][j] = p
						break
					}

				}
			}
		}
	}
	return s.sudoku
}

func initPlacement(sudoku [][]int) []*node {
	placement := make([]*node, 81)
	for i, row := range sudoku {
		for j, cell := range row {
			cellNode := &node{}
			if cell > 0 {
				cellNode.possible = []int{cell}
			}
			placement[i*9+j] = cellNode
		}
	}
	return placement
}

func isComplete(filled []bool) bool {
	for _, cell := range filled {
		if !cell {
			return false
		}
	}
	return true
}

func getPossibleByElimination(row, col int, sudoku [][]int) []int {
	possible := map[int]byte{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0}
	removeNumbersAlreadyInRow(row, sudoku, possible)
	removeNumbersAlreadyInCol(col, sudoku, possible)
	removeNumbersAlreadyInBlock(row, col, sudoku, possible)
	nums := make([]int, len(possible))
	c := 0
	for k := range possible {
		nums[c] = k
		c++
	}
	return nums
}

func isSingleOccurrenceOfNumber(row, col, num int, placement []*node) bool {
	return isSingleOccurrenceInRow(row, num, placement) ||
		isSingleOccurrenceInCol(col, num, placement) ||
		isSingleOccurrenceInBlock(row, col, num, placement)
}

func removeNumbersAlreadyInRow(row int, sudoku [][]int, possible map[int]byte) {
	for i := 0; i < 9; i++ {
		if sudoku[row][i] > 0 {
			delete(possible, sudoku[row][i])
		}
	}
}

func removeNumbersAlreadyInCol(col int, sudoku [][]int, possible map[int]byte) {
	for i := 0; i < 9; i++ {
		if sudoku[i][col] > 0 {
			delete(possible, sudoku[i][col])
		}
	}
}

func removeNumbersAlreadyInBlock(row, col int, sudoku [][]int, possible map[int]byte) {
	rowBlock := row % 3
	colBlock := col % 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := row + i - rowBlock
			c := col + j - colBlock
			if sudoku[r][c] > 0 {
				delete(possible, sudoku[r][c])
			}
		}
	}
}

func isSingleOccurrenceInRow(row, num int, placement []*node) bool {
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		cell := placement[row*9+i]
		if len(cell.possible) == 0 || contains(num, cell.possible) {
			return false
		}
	}
	return true
}

func isSingleOccurrenceInCol(col, num int, placement []*node) bool {
	for i := 0; i < 9; i++ {
		if i == col {
			continue
		}
		cell := placement[i*9+col]
		if len(cell.possible) == 0 || contains(num, cell.possible) {
			return false
		}
	}
	return true
}

func isSingleOccurrenceInBlock(row, col, num int, placement []*node) bool {
	rowBlock := row % 3
	colBlock := col % 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := row + i - rowBlock
			c := col + j - colBlock
			if r == row && c == col {
				continue
			}
			cell := placement[r*9+c]
			if len(cell.possible) == 0 || contains(num, cell.possible) {
				return false
			}
		}
	}
	return true
}

func contains(num int, array []int) bool {
	for _, n := range array {
		if n == num {
			return true
		}
	}
	return false
}

type node struct {
	possible []int
}
