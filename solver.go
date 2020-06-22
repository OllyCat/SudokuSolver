package main

import (
	"errors"
)

func (s *Sudoku) Solve() bool {
	x, y, err := s.nextFreeCell()

	//fmt.Println(s)

	if err != nil {
		if s.validateSolution() {
			return true
		}
		return false
	}

	for v := 1; v <= 9; v++ {
		s.m[y][x] = v
		if s.validateSolution() {
			if s.Solve() {
				return true
			}
		}
	}
	s.m[y][x] = 0
	return false
}

func (s *Sudoku) nextFreeCell() (int, int, error) {
	for i := 0; i < len(s.m); i++ {
		for j := 0; j < len(s.m[0]); j++ {
			if s.m[i][j] == 0 {
				return j, i, nil
			}
		}
	}
	return 0, 0, errors.New("Could not found free cell")
}
