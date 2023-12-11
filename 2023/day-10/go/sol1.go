package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

var (
	pipes = []byte{'|', '-', 'L', 'J', '7', 'F'}
)

func Max(arr ...int) int {
	max := arr[0]
	for _, x := range arr[1:] {
		if x > max {
			max = x
		}
	}
	return max
}

func checkPipeConnection(prev, curr byte, dir int) bool {
	var (
		invalidDown,
		invalidLeft,
		invalidRight,
		invalidUp bool
	)
	switch prev {
	case '|':
		invalidUp = dir == DOWN && curr != 'F' && curr != '7' && curr != '|'
		invalidDown = dir == UP && curr != 'L' && curr != 'J' && curr != '|'
	case '-':
		invalidLeft = dir == RIGHT && curr != 'F' && curr != 'L' && curr != '-'
		invalidRight = dir == LEFT && curr != 'J' && curr != '7' && curr != '-'
	case 'L':
		invalidUp = dir == LEFT && curr != 'J' && curr != '7' && curr != '-'
		invalidRight = dir == DOWN && curr != 'F' && curr != '7' && curr != '|'
	case 'J':
		invalidUp = dir == DOWN && curr != 'F' && curr != '7' && curr != '|'
		invalidLeft = dir == RIGHT && curr != 'F' && curr != 'L' && curr != '-'
	case '7':
		invalidDown = dir == UP && curr != 'L' && curr != 'J' && curr != '|'
		invalidLeft = dir == RIGHT && curr != 'F' && curr != 'L' && curr != '-'
	case 'F':
		invalidDown = dir == UP && curr != 'L' && curr != 'J' && curr != '|'
		invalidRight = dir == LEFT && curr != 'J' && curr != '7' && curr != '-'
	default:
		return false
	}
	if invalidDown || invalidLeft || invalidRight || invalidUp {
		return false
	}
	return true
}

func walk(grid [][]byte, visited [][]bool, prev byte, dir int, si, sj int, ei, ej int) (int, bool) {
	// fmt.Println(si, sj, string(grid[si][sj]), dir)
	if si < 0 || si >= len(grid) || sj < 0 || sj >= len(grid[0]) {
		return 0, false
	}

	curr := grid[si][sj]
	if visited[si][sj] || !checkPipeConnection(prev, curr, dir) {
		return 0, false
	}

	visited[si][sj] = true
	if si == ei && sj == ej {
		return 1, true
	}

	_si, _sj, _dir := si, sj, UP
	switch curr {
	case '|':
		if dir == UP {
			_si, _sj, _dir = si+1, sj, UP
		} else {
			_si, _sj, _dir = si-1, sj, DOWN
		}
	case '-':
		if dir == LEFT {
			_si, _sj, _dir = si, sj+1, LEFT
		} else {
			_si, _sj, _dir = si, sj-1, RIGHT
		}
	case 'L':
		if dir == UP {
			_si, _sj, _dir = si, sj+1, LEFT
		} else {
			_si, _sj, _dir = si-1, sj, DOWN
		}
	case 'J':
		if dir == UP {
			_si, _sj, _dir = si, sj-1, RIGHT
		} else {
			_si, _sj, _dir = si-1, sj, DOWN
		}
	case '7':
		if dir == DOWN {
			_si, _sj, _dir = si, sj-1, RIGHT
		} else {
			_si, _sj, _dir = si+1, sj, UP
		}
	case 'F':
		if dir == DOWN {
			_si, _sj, _dir = si, sj+1, LEFT
		} else {
			_si, _sj, _dir = si+1, sj, UP
		}
	}
	res, ok := walk(grid, visited, curr, _dir, _si, _sj, ei, ej)
	if !ok {
		return 0, false
	}
	return res + 1, true
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*20)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	grid := bytes.Split(buf, []byte("\n"))
	si, sj := 0, 0
	for i, line := range grid {
		if j := bytes.IndexByte(line, 'S'); j != -1 {
			si, sj = i, j
		}
	}
	ans, n, m := 0, len(grid), len(grid[0])
	for _, pipe := range pipes {
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, m)
		}
		grid[si][sj] = pipe
		res := 0
		switch pipe {
		case '|':
			res, _ = walk(grid, visited, '|', UP, si+1, sj, si, sj)
		case '-':
			res, _ = walk(grid, visited, '-', LEFT, si, sj+1, si, sj)
		case 'L':
			res, _ = walk(grid, visited, 'L', LEFT, si, sj+1, si, sj)
		case 'J':
			res, _ = walk(grid, visited, 'J', DOWN, si-1, sj, si, sj)
		case '7':
			res, _ = walk(grid, visited, '7', UP, si+1, sj, si, sj)
		case 'F':
			res, _ = walk(grid, visited, 'F', LEFT, si, sj+1, si, sj)
		}
		fmt.Println(pipe, res)
		ans = Max(ans, res)
	}
	fmt.Println(ans / 2)
}
