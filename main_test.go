package solver

import (
	"testing"
)

func Test_SolveByElimination(t *testing.T) {
	input := [][]int{
		{0, 0, 8, 2, 0, 0, 9, 0, 3},
		{3, 4, 2, 0, 9, 5, 0, 0, 7},
		{1, 9, 7, 0, 0, 0, 0, 0, 4},

		{0, 0, 5, 3, 1, 2, 4, 7, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 7, 4, 5, 0, 0},

		{0, 2, 0, 0, 0, 1, 0, 0, 5},
		{0, 7, 0, 0, 0, 6, 8, 9, 1},
		{8, 0, 0, 4, 3, 0, 7, 0, 6}}

	expected := [][]int{
		{5, 6, 8, 2, 4, 7, 9, 1, 3},
		{3, 4, 2, 1, 9, 5, 6, 8, 7},
		{1, 9, 7, 8, 6, 3, 2, 5, 4},
		{6, 8, 5, 3, 1, 2, 4, 7, 9},
		{7, 3, 4, 9, 5, 8, 1, 6, 2},
		{2, 1, 9, 6, 7, 4, 5, 3, 8},
		{9, 2, 6, 7, 8, 1, 3, 4, 5},
		{4, 7, 3, 5, 2, 6, 8, 9, 1},
		{8, 5, 1, 4, 3, 9, 7, 2, 6}}
	solver := NewSolver(input)
	solved := solver.Solve()

	for i, row := range expected {
		for j := range row {
			if expected[i][j] != solved[i][j] {
				t.Errorf("expected: %v, actal: %v", expected, solved)
			}
		}
	}
}

func Test_SolveByEliminationAndOccurrence(t *testing.T) {
	input := [][]int{
		{0, 0, 0, 0, 3, 0, 0, 0, 7},
		{0, 7, 0, 0, 0, 0, 1, 2, 0},
		{1, 0, 0, 0, 6, 4, 5, 8, 0},

		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		{5, 0, 0, 0, 0, 9, 7, 6, 0},
		{7, 4, 0, 0, 0, 0, 0, 1, 9},

		{0, 0, 8, 4, 2, 0, 0, 0, 1},
		{4, 0, 2, 0, 1, 0, 6, 7, 8},
		{0, 0, 0, 0, 0, 0, 0, 4, 0}}

	expected := [][]int{
		{8, 6, 5, 1, 3, 2, 4, 9, 7},
		{3, 7, 4, 5, 9, 8, 1, 2, 6},
		{1, 2, 9, 7, 6, 4, 5, 8, 3},

		{2, 9, 6, 8, 7, 1, 3, 5, 4},
		{5, 8, 1, 3, 4, 9, 7, 6, 2},
		{7, 4, 3, 2, 5, 6, 8, 1, 9},

		{6, 5, 8, 4, 2, 7, 9, 3, 1},
		{4, 3, 2, 9, 1, 5, 6, 7, 8},
		{9, 1, 7, 6, 8, 3, 2, 4, 5}}
	solver := NewSolver(input)
	solved := solver.Solve()

	for i, row := range expected {
		for j := range row {
			if expected[i][j] != solved[i][j] {
				t.Errorf("expected: %v, actal: %v", expected, solved)
			}
		}
	}
}
