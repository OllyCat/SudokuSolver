package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var item = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{5, 6, 7, 8, 9, 1, 2, 3, 4},
	{8, 9, 1, 2, 3, 4, 5, 6, 7},
	{3, 4, 5, 6, 7, 8, 9, 1, 2},
	{6, 7, 8, 9, 1, 2, 3, 4, 5},
	{9, 1, 2, 3, 4, 5, 6, 7, 8},
}

var withFree = [][]int{
	{8, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 3, 6, 0, 0, 0, 0, 0},
	{0, 7, 0, 0, 9, 0, 2, 0, 0},
	{0, 5, 0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 4, 5, 7, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 3, 0},
	{0, 0, 1, 0, 0, 0, 0, 6, 8},
	{0, 0, 8, 5, 0, 0, 0, 1, 0},
	{0, 9, 0, 0, 0, 0, 4, 0, 0},
}

var resolved = [][]int{
	{8, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 3, 6, 0, 0, 0, 0, 0},
	{0, 7, 0, 0, 9, 0, 2, 0, 0},
	{0, 5, 0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 4, 5, 7, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 3, 0},
	{0, 0, 1, 0, 0, 0, 0, 6, 8},
	{0, 0, 8, 5, 0, 0, 0, 1, 0},
	{0, 9, 0, 0, 0, 0, 4, 0, 0},
}

var itemBad = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{5, 6, 7, 8, 9, 1, 2, 3, 4},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{3, 4, 5, 6, 7, 8, 9, 1, 2},
	{6, 7, 8, 9, 1, 2, 3, 4, 5},
	{9, 1, 2, 3, 4, 5, 6, 7, 8},
}

var itemBadFree = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 0, 1, 2, 3, 4, 5, 6},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{5, 6, 7, 8, 0, 1, 2, 3, 4},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{3, 4, 5, 6, 7, 8, 0, 1, 2},
	{6, 7, 8, 9, 1, 2, 3, 4, 5},
	{9, 1, 2, 3, 4, 5, 6, 0, 8},
}

var itemBad2 = [][]int{
	{1, 2, 3, 4, 5, 3, 7, 8, 9},
	{4, 5, 6, 7, 8, 6, 1, 2, 3},
	{7, 8, 9, 1, 2, 9, 4, 5, 6},
	{2, 3, 4, 5, 6, 4, 8, 9, 1},
	{5, 6, 7, 8, 9, 7, 2, 3, 4},
	{8, 9, 1, 2, 3, 1, 5, 6, 7},
	{3, 4, 5, 6, 7, 5, 9, 1, 2},
	{6, 7, 8, 9, 1, 8, 3, 4, 5},
	{9, 1, 2, 3, 4, 2, 6, 7, 8},
}

func TestPrint(t *testing.T) {
	s := NewSudoku(item)
	fmt.Println(s)
}

func TestGetRow(t *testing.T) {
	s := NewSudoku(item)
	res := s.getRow(2)
	if !reflect.DeepEqual(res, []int{7, 8, 9, 1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}

func TestGetCol(t *testing.T) {
	s := NewSudoku(item)
	res := s.getCol(1)
	if !reflect.DeepEqual(res, []int{2, 5, 8, 3, 6, 9, 4, 7, 1}) {
		t.Fail()
	}
}

func TestGetBlock(t *testing.T) {
	s := NewSudoku(item)
	res := s.getBlock(1, 2)
	if !reflect.DeepEqual(res, [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}) {
		t.Fail()
	}

	res = s.getBlock(4, 7)
	if !reflect.DeepEqual(res, [][]int{
		{6, 7, 8},
		{9, 1, 2},
		{3, 4, 5},
	}) {
		t.Fail()
	}
}

func TestValidSlice(t *testing.T) {
	// valid slice
	s := NewSudoku(item)
	for i := 0; i < 9; i++ {
		res := s.getRow(i)
		if !validSlice(res) {
			t.Fail()
		}
	}
	// invalid slice
	s = NewSudoku(itemBad)
	for i := 0; i < 9; i++ {
		res := s.getCol(i)
		if validSlice(res) {
			t.Fail()
		}
	}
}

func TestValidBlock(t *testing.T) {
	// valid block
	s := NewSudoku(item)
	res := s.getBlock(3, 3)
	if !validBlock(res) {
		t.Fail()
	}
	// invalid block
	s = NewSudoku(itemBad)
	res = s.getBlock(3, 3)
	if validBlock(res) {
		t.Fail()
	}
}

func TestValidateSolution(t *testing.T) {
	s := NewSudoku(item)
	if !s.validateSolution() {
		t.Fatal("Get valid item, but got invalid ansver.")
	}
	s = NewSudoku(itemBad)
	if s.validateSolution() {
		t.Fatal("Get invalid item, but got valid ansver.")
	}
	s = NewSudoku(itemBad2)
	if s.validateSolution() {
		t.Fatal("Get invalid item, but got valid ansver.")
	}
	s = NewSudoku(withFree)
	if !s.validateSolution() {
		t.Fatal("Get valid item, but got invalid ansver.")
	}
}

func TestNextFreeCell(t *testing.T) {
	var res = [][]int{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
		{7, 0},
		{8, 0},
		{0, 1},
		{1, 1},
		{4, 1},
	}

	s := NewSudoku(withFree)
	for i := 0; i < 5; i++ {
		x, y, err := s.nextFreeCell()
		log.Printf("x = %d, y = %d, err = %v\n", x, y, err)
		if err != nil {
			t.Fatal("Has free cell, but has error.")
		}
		if x != res[i][0] || y != res[i][1] {
			t.Fatalf("Incorrect result: x = %d, y = %d\n", x, y)
		}
		s.m[y][x]++
	}
	s = NewSudoku(item)
	x, y, err := s.nextFreeCell()
	log.Printf("x = %d, y = %d, err = %v\n", x, y, err)
	if err == nil {
		t.Fatal("No free cells but has error")
	}
}

func TestSolver(t *testing.T) {
	s := NewSudoku(item)
	sr := NewSudoku(resolved)
	sb := NewSudoku(itemBad)

	res := s.Solve()
	log.Printf("Solvig full matrix: %v\n", res)
	if !res {
		t.Fail()
	}

	res = sr.Solve()
	log.Printf("Solvig resolved matrix: %v\n", res)
	if !res {
		t.Fail()
	}

	res = sb.Solve()
	log.Printf("Solvig unresolved matrix: %v\n", res)
	if res {
		t.Fail()
	}
}

func BenchmarkSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sr := NewSudoku(resolved)
		_ = sr.Solve()
	}
}
