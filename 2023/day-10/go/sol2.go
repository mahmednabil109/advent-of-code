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

var pipeDirs map[byte]byte = map[byte]byte{
	'|': 0b1100,
	'-': 0b0011,
	'L': 0b1010,
	'J': 0b1001,
	'7': 0b0101,
	'F': 0b0110,
}

func connected(grid [][]byte, i, j, k int) bool {
	m := len(grid[0])
	if j < 0 || j >= m || k < 0 || k >= m {
		return false
	}
	d1, d2 := pipeDirs[grid[i][j]], pipeDirs[grid[i][k]]
	return d1&0b0001 == 0b0001 && d2&0b0010 == 0b0010
}

func isInside(grid [][]byte, i, j int) bool {
	var (
		cs   byte
		cidx int
	)
	m, cuts, isEdge := len(grid[0]), 0, false

	for k := j; k < m; k++ {
		curr := grid[i][k]
		if isEdge {
			prev := grid[i][k-1]
			if curr == '.' {
				if cidx != k-1 && pipeDirs[cs]&pipeDirs[prev]&0b1100 != 0b0000 {
					cuts++
				}
				cuts++
				isEdge = false
			} else if !connected(grid, i, k, k-1) {
				if cidx != k-1 && pipeDirs[cs]&pipeDirs[prev]&0b1100 != 0b0000 {
					cuts++
				}
				cs = curr
				cidx = k
				cuts++
			}
		} else if curr != '.' {
			isEdge = true
			cs = curr
			cidx = k
		}
	}
	if isEdge {
		if cidx != m-1 && pipeDirs[cs]&pipeDirs[grid[i][m-1]]&0b1100 != 0b0000 {
			cuts++
		}
		cuts++
	}
	return cuts%2 == 1
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
	largestLoop := make([][]bool, 0)
	startPipe := byte(0)

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
		if res > ans {
			ans = res
			largestLoop = visited
			startPipe = pipe
		}
	}

	grid[si][sj] = startPipe
	for i, line := range largestLoop {
		for j, v := range line {
			if !v {
				grid[i][j] = '.'
			}
		}
	}

	ans = 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if !largestLoop[i][j] && grid[i][j] == '.' && isInside(grid, i, j) {
				ans++
				grid[i][j] = 'I'
			}
		}
	}

	fmt.Println(ans)
	for _, line := range grid {
		fmt.Println(string(line))
	}

}
