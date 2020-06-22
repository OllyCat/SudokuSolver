package main

import "fmt"

// структура поля и её методы
type Sudoku struct {
	m [][]int
}

// Создание нового класса Sudoku
func NewSudoku(m [][]int) *Sudoku {
	s := new(Sudoku)
	s.m = m
	return s
}

// Проверка всего поля.
// Входные: полне
// Выходные: ложь или истина
func (s *Sudoku) validateSolution() (res bool) {
	res = true
	for i := 0; i < len(s.m); i++ {
		if !validSlice(s.getRow(i)) {
			res = false
			break
		}
		if !validSlice(s.getCol(i)) {
			res = false
			break
		}
		y := i / 3 * 3
		x := i % 3 * 3
		if !validBlock(s.getBlock(x, y)) {
			res = false
			break
		}
	}
	return
}

// Получение строки.
// Входные: поле, y координата.
// Выходные: слайс содержащий строку
func (s *Sudoku) getRow(y int) []int {
	return s.m[y]
}

// Получить колонку.
// Входные параметры: поле, х координата.
// Выходные: слайс содержащий колонку
func (s *Sudoku) getCol(x int) []int {
	c := make([]int, len(s.m))
	for k, v := range s.m {
		c[k] = v[x]
	}
	return c
}

// Получить блок.
// Входные: поле, x и y координаты любой точки на поле
// Выходные: квадратное поле, в котором находится данная точка
func (s *Sudoku) getBlock(x, y int) [][]int {
	x = x / 3 * 3
	y = y / 3 * 3
	res := make([][]int, 3)

	for i := 0; i < 3; i++ {
		res[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i][j] = s.m[i+y][j+x]
		}
	}

	return res
}

// Проверка валидности слайса. Слайс не должен содержать одинаковых цифр, кроме нуля (нули - пустые места в случае нерешённого поля)
func validSlice(s []int) bool {
	// создаём словарь значений
	d := make(map[int]int, len(s))

	// проходим по всему массиву считая каждое число, встреченное в массиве
	for _, v := range s {
		d[v]++
	}

	// проверяем нет ли повторов любых цифр, кроме нуля
	for k, v := range d {
		if k != 0 && v > 1 {
			return false
		}
	}
	return true
}

// Проверка валидности блока.
// Входные: блок 3x3
// Выходные: истина, если блок валидный, ложь, если нет.
func validBlock(bl [][]int) bool {
	s := make([]int, len(bl)*len(bl[0]))
	var i int
	for _, v := range bl {
		for _, j := range v {
			s[i] = j
			i++
		}
	}
	return validSlice(s)
}

func (s *Sudoku) String() string {
	var res string
	res = "┌───────┬───────┬───────┐\n"
	for k, v := range s.m {
		res = res + "│"
		for i := 0; i < len(v); i += 3 {
			res = res + fmt.Sprintf(" %d %d %d │", v[i], v[i+1], v[i+2])
		}
		res = res + "\n"
		if k == 2 || k == 5 {
			res = res + "├───────┼───────┼───────┤\n"
		}
	}
	res = res + "└───────┴───────┴───────┘"
	return res
}
